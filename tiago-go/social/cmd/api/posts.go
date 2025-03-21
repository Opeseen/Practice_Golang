package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ope/social/internal/store"
)

// Handler to create a new post
func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	type CreatePostPayload struct {
		Title   string   `json:"title" validate:"required,max=100"`
		Content string   `json:"content" validate:"required,max=100"`
		Tags    []string `json:"tags"`
	}
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	// validate the create post payload
	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		// @TODO Change after auth
		UserID: 1,
	}
	ctx := r.Context()
	// this is based on the create method calling
	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	// this is based on the return of the create method calling
	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Handler to get a post
func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	ctx := r.Context()

	post, err := app.store.Posts.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundError(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	comments, err := app.store.Comments.GetPostByID(ctx, id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	post.Comments = comments // append any comment to the returned post

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
