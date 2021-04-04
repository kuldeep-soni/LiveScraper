package document_parser

import (
	"context"
	"errors"
	"fmt"
	"github.com/LiveScraper/app/website-scraping/scraper/model"
)

type IDocumentParserFactory interface {
	GetParser(ctx context.Context, source model.Source, parserType string) (parser IDocumentParser, err error)
}

type IDocumentParser interface {
	TransformRawMovieData(ctx context.Context, rawData string) (meta model.MovieMeta, err error)
}

type documentParserFactory struct{}

func (d *documentParserFactory) GetParser(ctx context.Context, source model.Source, parserType string) (parser IDocumentParser, err error) {
	switch getUniqueParserName(source, parserType) {
	case getUniqueParserName(model.Amazon, "mock"):
		return &amazonMockDocumentParser{}, nil
	case getUniqueParserName(model.Amazon, "type-1"):
		return &amazonDocumentParser{}, nil
	}
	return nil, errors.New(fmt.Sprintf("No parser found with source: %v and parserType: %v", source, parserType))
}

func getUniqueParserName(source model.Source, parserType string) string {
	return fmt.Sprintf("%v_%v", source, parserType)
}

func NewDocumentParserFactory() IDocumentParserFactory {
	return &documentParserFactory{}
}
