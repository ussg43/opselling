package db

import (
	"context"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ussg43/opselling/internal/env"
)

var (
	connInfo string = "ADD LATER"
)

func New() (*sqlx.DB, error) {
	dbConn := env.GetEnvVar("DB_CONNECTION", connInfo)
	db, err := sqlx.Connect("pgx", dbConn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	log.Printf("Successfully connected to:%s", db.DriverName())
	return db, nil
}
