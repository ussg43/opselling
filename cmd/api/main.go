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
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func main() {
	httpPort := env.GetEnvVar("HTTP_PORT", "5656")
	grpcPort := env.GetEnvVar("RPC_PORT", "135")
	configFile := env.GetEnvVar("OAUTH_CONFIG", "")

	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("could not load oauth client secret %v", err)
	}

	oauthcfg, err := google.ConfigFromJSON(file, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("%v", err)
	}

	sheetsConfig := &sheetsConfig{
		oauthConfig: oauthcfg,
	}

	config := &config{
		httpPort: httpPort,
		grpcPort:  grpcPort,
		sheetsConfig: *sheetsConfig,
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
