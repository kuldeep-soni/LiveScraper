package streaming_services

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/scraper/streaming-services/parsers"
	"github.com/LiveScraper/models"
)

type IStreamingServiceFactory interface {
	SetStreamingServices(streamingServices models.StreamingServices) (err error)
	GetStreamingService(ctx context.Context, name StreamingServiceName) IStreamingService
}

type streamingServiceFactory struct {
	config        map[StreamingServiceName]IStreamingService
	parserFactory parsers.IDocumentParserFactory
}

func (s *streamingServiceFactory) SetStreamingServices(streamingServices models.StreamingServices) (err error) {
	s.config = make(map[StreamingServiceName]IStreamingService)
	for _, ss := range streamingServices {
		sourceName := s.getStreamingServiceName(ss.Name)
		parserType := s.parserFactory.GetParserType(ss.ParserType)
		parser, err := s.parserFactory.GetParser(context.Background(), parserType)
		if err != nil {
			return err
		}

		if sourceName != "" && parserType != "" {
			s.config[sourceName] = &streamingService{
				Name:           sourceName,
				URL:            ss.URL,
				ParserType:     parserType,
				documentParser: parser,
			}
		}
	}
	return
}

func (s *streamingServiceFactory) GetStreamingService(ctx context.Context, name StreamingServiceName) (src IStreamingService) {
	return s.config[name]
}

func (s *streamingServiceFactory) getStreamingServiceName(name string) (srcName StreamingServiceName) {
	switch name {
	case string(Amazon):
		return Amazon
	}
	return
}

func NewStreamingServiceFactory(documentParserFactory parsers.IDocumentParserFactory) IStreamingServiceFactory {
	return &streamingServiceFactory{
		parserFactory: documentParserFactory,
	}
}
