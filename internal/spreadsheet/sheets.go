package spreadsheet

import (
	"fmt"
	"log"
	"time"

	"github.com/ussg43/opselling/internal/scraper"
	"google.golang.org/api/sheets/v4"
)


type RowData struct {
	Name string
	Id int64
	Price int64
	LastUpdated time.Time
}



func ReadSheet(srv *sheets.Service, ID string) [][]string{
	data := make([][]any,0)
	readRange := "Sheet1!A2:B3"
	res, err := srv.Spreadsheets.Values.Get(ID, readRange).Do()
	if err != nil{
		fmt.Errorf("Unable to retrieve data from sheet: %v", err)
		return nil
	}

	if len(res.Values) == 0{
		fmt.Println("no data")
	}else{
		data = append(data, res.Values...)
	}

	rows := make([][]string, 0)
	for _, val := range data{
		row := []string {val[0].(string), val[1].(string), val[2].(string)}
		rows = append(rows, row)
	}
	return rows
}

func WritetoSheet(srv *sheets.Service, sheetID string, data [][]string){
	writeRange := "Sheet1!A2"
	newData := make([][]any, 0)

	for i, row := range data{
		time.Sleep(10 * time.Second)
		price, err := scraper.GetCardData(row[0],row[1])
		if err != nil{
			fmt.Errorf("Unable to scrape player data: %v", err)
			return
		}else{
			row = append(row, price)
			fmt.Println(row)
			newData[i] = append(newData[i], row)
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