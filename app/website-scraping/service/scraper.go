package service

import (
	"context"
	"fmt"
	"github.com/LiveScraper/app/website-scraping/enums"
	"github.com/LiveScraper/app/website-scraping/model"
	"github.com/LiveScraper/app/website-scraping/streaming-services"
	"github.com/LiveScraper/phttp/client/httpClient"
)

//This interface is used by scraper transport layer
type IScraper interface {
	GetMovieMeta(ctx context.Context, sourceName enums.StreamingServiceName, movieId string) (model.MovieMeta, error)
}

//This is the scraper service which fetches documents from streaming service endpoints
//via networkConnector and coordinates with various streaming services via streamingServiceFactory
type scraper struct {
	networkConnector        httpClient.IHttpClient
	streamingServiceFactory streaming_services.IStreamingServiceFactory
}

func (s *scraper) GetMovieMeta(ctx context.Context, streamingServiceName enums.StreamingServiceName, movieId string) (meta model.MovieMeta, err error) {
	streamingService, err := s.streamingServiceFactory.GetStreamingService(ctx, streamingServiceName)
	if err != nil{
		return
	}
	websiteUrl := streamingService.GetUrl()

	resource, err := s.networkConnector.GetHTMLResource(ctx, fmt.Sprintf("%v/%v", websiteUrl, movieId))
	if err != nil {
		return
	}

	return streamingService.GetMovieMeta(ctx, resource)
}

//Initialises scraper service
func NewScraper(networkConnector httpClient.IHttpClient, streamingServiceFactory streaming_services.IStreamingServiceFactory) IScraper {
	return &scraper{
		networkConnector:        networkConnector,
		streamingServiceFactory: streamingServiceFactory,
	}
}
