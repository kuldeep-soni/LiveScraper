package streaming_services

import (
	"context"
	"errors"
	"fmt"
	"github.com/LiveScraper/app/website-scraping/enums"
	"github.com/LiveScraper/app/website-scraping/streaming-services/amazon"
	"github.com/LiveScraper/global"
)

//This factory initialises all streaming services as well as fetches a streaming service by its name
//It is used by scraper service
type IStreamingServiceFactory interface {
	SetStreamingServices(streamingServices global.StreamingServices) (err error)
	GetStreamingService(ctx context.Context, name enums.StreamingServiceName) (IStreamingService, error)
}

type streamingServiceFactory struct {
	config map[enums.StreamingServiceName]IStreamingService
}

func (s *streamingServiceFactory) SetStreamingServices(streamingServices global.StreamingServices) (err error) {
	s.config = make(map[enums.StreamingServiceName]IStreamingService)
	for _, ss := range streamingServices {
		sourceName, err := s.getStreamingServiceName(ss.Name)
		if err != nil {
			return err
		}

		switch sourceName {
		case enums.Amazon:
			s.config[sourceName], err = amazon.NewAmazonStreamingService(sourceName, ss.URL, ss.ParserType)
		}
		if err != nil {
			return err
		}
	}
	return
}

func (s *streamingServiceFactory) GetStreamingService(ctx context.Context, name enums.StreamingServiceName) (src IStreamingService, err error) {
	src = s.config[name]
	if src == nil || src.GetName() == "" {
		return nil, errors.New(fmt.Sprintf("%v streaming service is not configured", name))
	}
	return
}

func (s *streamingServiceFactory) getStreamingServiceName(name string) (srcName enums.StreamingServiceName, err error) {
	switch name {
	case string(enums.Amazon):
		return enums.Amazon, nil
	}
	return srcName, errors.New(fmt.Sprintf("%v is not a valid streaming service", name))
}

func NewStreamingServiceFactory() IStreamingServiceFactory {
	return &streamingServiceFactory{}
}
