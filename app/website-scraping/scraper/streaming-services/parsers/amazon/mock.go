package amazon

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/scraper/model"
)

type amazonMockDocumentParser struct {
}

func GetAmazonMockDocumentParser() *amazonMockDocumentParser {
	return &amazonMockDocumentParser{}
}

func (am *amazonMockDocumentParser) TransformRawMovieData(ctx context.Context, rawData string) (meta model.MovieMeta, err error) {
	meta.Title = "Test"
	meta.ReleaseYear = 2021
	meta.Actors = []string{"Kuldeep Soni"}
	meta.SimilarIds = []string{"K", "U", "L"}
	return
}
