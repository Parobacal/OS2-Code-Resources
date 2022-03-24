package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("could not create the file, err :%q", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)
	c.OnHTML(".internship_meta", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("a"),
		})
	})

	for i := 0; i < 330; i++ {
		fmt.Printf("Scraping Page : %d\n", i)
		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}
	log.Printf("Scraping Complete\n")
	log.Println(c)
}
