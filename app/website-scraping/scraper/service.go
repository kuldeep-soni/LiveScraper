package scraper

import (
	"context"
	"fmt"
	"github.com/LiveScraper/app/website-scraping/scraper/model"
	"github.com/LiveScraper/phttp/client/httpClient"
)

type IService interface {
	GetMovieMeta(ctx context.Context, movieId string) (model.MovieMeta, error)
}

type service struct {
	networkConnector httpClient.IHttpClient
	documentReader   IDocumentReader
}

func (s *service) GetMovieMeta(ctx context.Context, movieId string) (meta model.MovieMeta, err error) {
	resource, err := s.networkConnector.GetHTMLResource(ctx, fmt.Sprintf("http://www.amazon.de/gp/product/%v", movieId))
	if err != nil {
		return
	}

	return s.documentReader.TransformRawMovieData(ctx, resource)
}

func NewService(httpClient httpClient.IHttpClient, documentReader IDocumentReader) IService {
	return &service{
		networkConnector: httpClient,
		documentReader:   documentReader,
	}
}
