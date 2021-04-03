package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")
}

//extendable
//web server
//scraping service -> scrape based on which website to scrape, which location to scrape, have a config file which can define
//all the css selectors so that redeploy can just suffice
//what if new data is to be included?
//including a service which escapes