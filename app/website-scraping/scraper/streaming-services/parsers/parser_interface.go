package parsers

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/scraper/model"
)

type ParserType string

const (
	AmazonParser1    = ParserType("amazon-parser-1")
	AmazonMockParser = ParserType("amazon-mock-parser")
)

type IDocumentParser interface {
	TransformRawMovieData(ctx context.Context, rawData string) (meta model.MovieMeta, err error)
}
