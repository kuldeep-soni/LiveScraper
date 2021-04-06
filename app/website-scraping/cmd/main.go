package main

import (
	"github.com/LiveScraper/app/website-scraping/cmd/commands"
	"os"
)

//This is the main program. Run this to start the website-scraping application
func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
