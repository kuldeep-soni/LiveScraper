package scraper

import (
	"context"
	"fmt"
	"github.com/LiveScraper/app/website-scraping/scraper/model"
	"github.com/LiveScraper/app/website-scraping/scraper/streaming-services"
	"github.com/LiveScraper/phttp/client/httpClient"
)

type IService interface {
	GetMovieMeta(ctx context.Context, sourceName streaming_services.StreamingServiceName, movieId string) (model.MovieMeta, error)
}

type service struct {
	networkConnector        httpClient.IHttpClient
	streamingServiceFactory streaming_services.IStreamingServiceFactory
}

func (s *service) GetMovieMeta(ctx context.Context, streamingServiceName streaming_services.StreamingServiceName, movieId string) (meta model.MovieMeta, err error) {
	streamingService := s.streamingServiceFactory.GetStreamingService(ctx, streamingServiceName)
	websiteUrl := streamingService.GetUrl()

	resource, err := s.networkConnector.GetHTMLResource(ctx, fmt.Sprintf("%v/%v", websiteUrl, movieId))
	if err != nil {
		return
	}

	return streamingService.GetMovieMeta(ctx, resource)
}

func NewService(networkConnector httpClient.IHttpClient, streamingServiceFactory streaming_services.IStreamingServiceFactory) IService {
	return &service{
		networkConnector:        networkConnector,
		streamingServiceFactory: streamingServiceFactory,
	}
}
