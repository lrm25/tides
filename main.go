package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	collector := colly.NewCollector()
	collector.OnHTML(".mt-1.font-weight-normal", func(element *colly.HTMLElement) {
		fmt.Println(element.Text)
	})
	if err := collector.Visit("https://www.tideschart.com/Brazil/Paraiba/Cabedelo/"); err != nil {
		fmt.Println(err.Error())
	}
}
