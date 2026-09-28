package main

import (
	"log"

	"github.com/ussg43/opselling/internal/db"
	"github.com/ussg43/opselling/internal/env"
	"github.com/ussg43/opselling/internal/store"
)

func main() {
	conStr := env.GetEnvVar("DB_CONN", "str")
	httpPort := env.GetEnvVar("HTTP_PORT", "5656")
	rpcPort := env.GetEnvVar("RPC_PORT", "135")
	config := &config{
		dbConn:   conStr,
		httpPort: httpPort,
		rpcPort:  rpcPort,
	}

	db, err := db.New()
	if err != nil {
		log.Fatalf("Failed to establish db connection, %v", err)
	}

	storage := store.NewStorage(db)

	app := &application{
		config:  *config,
		db:      db,
		storage: storage,
	}

	app.mount()

}
