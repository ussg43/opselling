package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ussg43/opselling/internal/spreadsheet"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	sheetID := "1rf9JHoq1HItqsxCEYPXK5JKl_Or8KbOrfzw2yDmu2Js"
	ctx := context.Background()
	b, err := os.ReadFile("stable-liberty-474404-m1-5dd405f1795c.json")
	if err != nil {
		log.Fatalf("Cannot read client file")
	}

	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil{
		log.Fatalf("Unable to parse client file to config")
	}

	client := config.Client(ctx)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil{
		log.Fatalf("Unable to retrieve sheets client")
	}

	for{
		data := spreadsheet.ReadSheet(srv, sheetID)
		spreadsheet.WritetoSheet(srv, sheetID, data)
		time.Sleep(45 * time.Minute)
	}
}
