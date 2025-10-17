package spreadsheet

import (
	"fmt"
	"log"
	"time"

	"github.com/ussg43/opselling/internal/scraper"
	"google.golang.org/api/sheets/v4"
)



func ReadSheet(srv *sheets.Service, ID string) [][]any{
	data := make([][]any,0)
	readRange := "Sheet1!A2:B3"
	res, err := srv.Spreadsheets.Values.Get(ID, readRange).Do()
	if err != nil{
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(res.Values) == 0{
		fmt.Println("no data")
	}else{
		data = append(data, res.Values...)
	}
	return data
}

func WritetoSheet(srv *sheets.Service, sheetID string, data [][]any){
	writeRange := "Sheet1!A2"
	newData := make([][]any, 0)
	for _, row := range data{
		time.Sleep(10 * time.Second)
		price, err := scraper.GetCardData(row[0].(string),row[1].(string))
		if err != nil{
			log.Fatalf("Unable to scrape player data: %v", err)
		}else{
			row = append(row, price)
			fmt.Println(row)
			newData = append(newData, row)
		}
	}

	valueRange := &sheets.ValueRange{
		Values: newData,
		MajorDimension: "ROWS",
	}
	_, err := srv.Spreadsheets.Values.Update(sheetID,writeRange,valueRange).ValueInputOption("RAW").Do()
	if err != nil{
		log.Fatalf("Unable to write to sheet %v", err)
	}
}