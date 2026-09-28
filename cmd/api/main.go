package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ussg43/opselling/internal/db"
	"github.com/ussg43/opselling/internal/env"
	"github.com/ussg43/opselling/internal/store"
)

func main() {
	httpPort := env.GetEnvVar("HTTP_PORT", "5656")
	grpcPort := env.GetEnvVar("RPC_PORT", "135")
	config := &config{
		httpPort: httpPort,
		grpcPort:  grpcPort,
	}

	db, err := db.New()
	if err != nil {
		log.Fatalf("Failed to establish db connection, %v", err)
	}

	storage := store.NewStorage(db)

	app := &Application{
		config:  *config,
		db:      db,
		storage: storage,
	}

	app.mount()
	
	ctx, signal := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer signal()
	if err := app.run(ctx); err != nil {
		log.Fatal(err)
	}

}
