package parsers

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/enums"
	amazonModels "github.com/LiveScraper/app/website-scraping/streaming-services/amazon/models"
)

type amazonMockDocumentParser struct {
	name enums.ParserType
}

func (am *amazonMockDocumentParser) GetName() enums.ParserType {
	return am.name
}

func (am *amazonMockDocumentParser) TransformRawMovieData(ctx context.Context, rawData string) (meta amazonModels.MovieMetaAmazon, err error) {
	meta.Title = "Test"
	meta.ReleaseYear = 2021
	meta.Actors = []string{"Kuldeep Soni"}
	meta.SimilarIds = []string{"K", "U", "L"}
	return
}

func GetAmazonMockDocumentParser(parserType enums.ParserType) *amazonMockDocumentParser {
	return &amazonMockDocumentParser{name: parserType}
}