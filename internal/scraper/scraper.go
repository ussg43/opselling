package scraper

import (
	// "errors"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Player struct{
	Name string
	Price string
}

func GetCardData(name string, ID string) (string, error){
	price := ""
	name = strings.ReplaceAll(name, " ","-")
	url := strings.ToLower(fmt.Sprintf("https://futbin.com/26/player/%s/%s",ID,name))
	log.Println(url)
	c:= colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.AllowedDomains("www.futbin.com","futbin.com"),
	)

	c.OnHTML(".price.inline-with-icon.lowest-price-1", func(h *colly.HTMLElement){
		if price == ""{
			price = h.Text
		}
	})


	c.OnRequest(func(r *colly.Request){
		log.Println("visiting:", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error){
		log.Println("an error occured", err)
	})

	err:= c.Visit(url)

	if err != nil{
		return "", err
	}

	return price, nil
}