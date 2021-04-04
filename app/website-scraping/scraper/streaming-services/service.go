package streaming_services

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/scraper/model"
	"github.com/LiveScraper/app/website-scraping/scraper/streaming-services/parsers"
)

type StreamingServiceName string

const (
	Amazon = StreamingServiceName("amazon")
)

type IStreamingService interface {
	GetName() StreamingServiceName
	GetUrl() string
	GetMovieMeta(ctx context.Context, rawMovieData string) (meta model.MovieMeta, err error)
}

type streamingService struct {
	Name           StreamingServiceName
	URL            string
	ParserType     parsers.ParserType
	documentParser parsers.IDocumentParser
}

func (s *streamingService) GetName() StreamingServiceName {
	return s.Name
}

func (s *streamingService) GetUrl() string {
	return s.URL
}

func (s *streamingService) GetMovieMeta(ctx context.Context, rawMovieData string) (meta model.MovieMeta, err error) {
	return s.documentParser.TransformRawMovieData(ctx, rawMovieData)
}