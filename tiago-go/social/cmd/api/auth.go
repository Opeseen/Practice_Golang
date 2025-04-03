package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/ope/social/internal/mailer"
	"github.com/ope/social/internal/store"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
}

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	// Validate the payload
	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	// map payload to the user
	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	// hash the user password
	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	// generate a token
	plainToken := uuid.New().String()

	// hash the token
	hash := sha256.Sum256([]byte(plainToken))
	// convert the hash to a string
	hashToken := hex.EncodeToString(hash[:])

	// Create the new user
	err := app.store.Users.CreateAndInvite(ctx, user, hashToken, app.config.mail.exp)
	if err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			app.badRequestError(w, r, err)
		case store.ErrDuplicateUsername:
			app.badRequestError(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	userWithToken := UserWithToken{
		User:  user,
		Token: plainToken,
	}

	// Return JSON response immediately
	if err := app.jsonResponse(w, http.StatusCreated, userWithToken); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// ---- Start sending email notification asynchronously to the user  ----
	go func(user *store.User, plainToken string) {
		activationURL := fmt.Sprintf("%s/confirm/%s", app.config.frontendURL, plainToken)

		isProdEnv := app.config.env == "production"
		vars := struct {
			Username      string
			ActivationURL string
		}{
			Username:      user.Username,
			ActivationURL: activationURL,
		}

		// Send mail
		status, err := app.mailer.Send(mailer.UserWelcomeTemplate, user.Username, user.Email, vars, !isProdEnv)
		if err != nil {
			app.logger.Errorw("error sending welcome email", "error", err)

			// create another background context
			bgCtx := context.Background()
			// Rollback user creation if email fails (SAGA pattern)
			if err := app.store.Users.Delete(bgCtx, user.ID); err != nil {
				app.logger.Errorw("error deleting user", "error", err)
			}
			return
		}

		app.logger.Infow("Email sent", "status code", status)
	}(user, plainToken)
}
