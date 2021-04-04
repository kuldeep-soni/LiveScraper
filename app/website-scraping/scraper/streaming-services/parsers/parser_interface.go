package parsers

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/scraper/model"
)

type IDocumentParser interface {
	TransformRawMovieData(ctx context.Context, rawData string) (meta model.MovieMeta, err error)
}
