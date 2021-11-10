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
	fname := "data.csv"
	file, err := os.Create(fname)
	if err != nil {
		log.Fatal("Failed to create file, error: %q", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("tuoitre.vn"),
	)
	c.OnHTML(".name-news", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("a"),
			e.ChildText("p"),
			e.ChildText("span"),
		})
	})

	for i := 0; i < 100; i++ {
		fmt.Printf("Scraping page: %d\n", i)
		c.Visit("https://tuoitre.vn/kinh-doanh/doanh-nghiep/trang-" + strconv.Itoa(i) + ".htm")
	}

	log.Printf("Scraping Done\n")
	log.Println(c)
}
