package amazon

import (
	"context"
	"errors"
	"fmt"
	"github.com/LiveScraper/app/website-scraping/enums"
	"github.com/LiveScraper/app/website-scraping/streaming-services/amazon/parsers"
)

type IAmazonDocumentParserFactory interface {
	GetParserType(parserType string) (pt enums.ParserType, err error)
	GetParser(ctx context.Context, parserType enums.ParserType) (parser IAmazonDocumentParser, err error)
}

type amazonDocumentParserFactory struct{}

func (d *amazonDocumentParserFactory) GetParserType(parserType string) (pt enums.ParserType, err error) {
	switch parserType {
	case string(enums.AmazonParser1):
		return enums.AmazonParser1, nil
	case string(enums.AmazonMockParser):
		return enums.AmazonMockParser, nil
	default:
		return pt, errors.New(fmt.Sprintf("%v is not a valid document parser for amazon streaming service", parserType))
	}
}

func (d *amazonDocumentParserFactory) GetParser(ctx context.Context, parserType enums.ParserType) (parser IAmazonDocumentParser, err error) {
	switch parserType {
	case enums.AmazonMockParser:
		return parsers.GetAmazonMockDocumentParser(parserType), nil
	case enums.AmazonParser1:
		return parsers.GetAmazonDocumentParser1(parserType), nil
	}
	return nil, errors.New(fmt.Sprintf("No parser found with parserType: %v for amazon streaming service", parserType))
}