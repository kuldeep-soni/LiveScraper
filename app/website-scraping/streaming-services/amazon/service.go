package amazon

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/enums"
	"github.com/LiveScraper/app/website-scraping/model"
)

type amazonStreamingService struct {
	Name                  enums.StreamingServiceName
	URL                   string
	ParserType            enums.ParserType
	documentParserFactory IAmazonDocumentParserFactory
}

func (s *amazonStreamingService) GetName() enums.StreamingServiceName {
	return s.Name
}

func (s *amazonStreamingService) GetUrl() string {
	return s.URL
}

func (s *amazonStreamingService) GetMovieMeta(ctx context.Context, rawMovieData string) (meta model.MovieMeta, err error) {
	parser, err := s.documentParserFactory.GetParser(ctx, s.ParserType)
	if err != nil {
		return
	}

	amazonMovieMeta, err := parser.TransformRawMovieData(ctx, rawMovieData)
	if err != nil {
		return
	}

	return amazonMovieMeta.GetGlobalMovieMeta(), nil
}

func NewAmazonStreamingService(name enums.StreamingServiceName, url string, parserType string) (*amazonStreamingService, error) {
	ss := &amazonStreamingService{}
	ss.documentParserFactory = &amazonDocumentParserFactory{}

	pt, err := ss.documentParserFactory.GetParserType(parserType)
	if err != nil {
		return nil, err
	}

	ss.ParserType = pt
	ss.Name = name
	ss.URL = url
	return ss, nil
}