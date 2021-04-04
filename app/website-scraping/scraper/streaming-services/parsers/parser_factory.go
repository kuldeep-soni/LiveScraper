package parsers

import (
	"context"
	"errors"
	"fmt"
	"github.com/LiveScraper/app/website-scraping/scraper/streaming-services/parsers/amazon"
)

type IDocumentParserFactory interface {
	GetParserType(parserType string) (pt ParserType)
	GetParser(ctx context.Context,  parserType ParserType) (parser IDocumentParser, err error)
}

type documentParserFactory struct{}

func (d *documentParserFactory) GetParserType(parserType string) (pt ParserType) {
	switch parserType {
	case string(AmazonParser1):
		return AmazonParser1
	case string(AmazonMockParser):
		return AmazonMockParser
	default:
		return
	}
}

func (d *documentParserFactory) GetParser(ctx context.Context, parserType ParserType) (parser IDocumentParser, err error) {
	switch parserType {
	case AmazonMockParser:
		return amazon.GetAmazonMockDocumentParser(), nil
	case AmazonParser1:
		return amazon.GetAmazonDocumentParser1(), nil
	}
	return nil, errors.New(fmt.Sprintf("No parser found with parserType: %v", parserType))
}

func NewDocumentParserFactory() IDocumentParserFactory {
	return &documentParserFactory{}
}
