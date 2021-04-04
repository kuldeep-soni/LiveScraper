package scraper

import (
	"context"
	"fmt"
	document_parser "github.com/LiveScraper/app/website-scraping/scraper/document-parser"
	"github.com/LiveScraper/app/website-scraping/scraper/model"
	"github.com/LiveScraper/phttp/client/httpClient"
)

type IService interface {
	GetMovieMeta(ctx context.Context, source model.Source, movieId string) (model.MovieMeta, error)
}

type service struct {
	networkConnector      httpClient.IHttpClient
	documentParserFactory document_parser.IDocumentParserFactory
}

func (s *service) GetMovieMeta(ctx context.Context, source model.Source, movieId string) (meta model.MovieMeta, err error) {
	websiteUrl, err := source.GetWebsiteUrl()
	if err != nil {
		return
	}

	resource, err := s.networkConnector.GetHTMLResource(ctx, fmt.Sprintf("%v/%v", websiteUrl, movieId))
	if err != nil {
		return
	}

	parser, err := s.documentParserFactory.GetParser(ctx, source, "type-1")
	if err != nil {
		return
	}

	return parser.TransformRawMovieData(ctx, resource)
}

func NewService(httpClient httpClient.IHttpClient, documentParserFactory document_parser.IDocumentParserFactory) IService {
	return &service{
		networkConnector:      httpClient,
		documentParserFactory: document_parser.NewDocumentParserFactory(),
	}
}
