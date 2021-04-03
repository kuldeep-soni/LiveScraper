package main

import (
	"github.com/LiveScraper/app/website-scraping/cmd/commands"
	"os"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
