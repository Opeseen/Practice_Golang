package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/ope/social/internal/db"
	"github.com/ope/social/internal/env"
	"github.com/ope/social/internal/mailer"
	"github.com/ope/social/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.1"

func main() {
	// Get the env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr:        env.GetString("ADDR", ":9090"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:5173"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:adminuser@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp:       time.Hour * 24 * 3, // 3days
			fromEmail: env.GetString("FROM_EMAIL", ""),
			mailTrap: mailTrapConfig{
				apiKey:      env.GetString("MAILTRAP_API_KEY", ""),
				apiUsername: env.GetString("MAILTRAP_USERNAME", ""),
			},
		},
	}

	// Zap Logger
	zapConfig := zap.NewDevelopmentConfig()
	// Disable stack trace for error logs
	zapConfig.DisableStacktrace = true
	logger := zap.Must(zapConfig.Build()).Sugar()
	defer logger.Sync()

	// Database
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}

	// close the db connection
	defer db.Close()
	logger.Info("Database connection pool established")

	mailtrap, err := mailer.NewMailTrapClient(cfg.mail.mailTrap.apiKey, cfg.mail.fromEmail, cfg.mail.mailTrap.apiUsername)
	if err != nil {
		logger.Fatal(err)
	}

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
		mailer: mailtrap,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
