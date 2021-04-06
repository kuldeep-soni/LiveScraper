package streaming_services

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/enums"
	"github.com/LiveScraper/app/website-scraping/model"
)

//This is a streaming service interface all streaming services must implement.
//scraper service only knows of IStreamingService
type IStreamingService interface {
	GetName() enums.StreamingServiceName
	GetUrl() string
	GetMovieMeta(ctx context.Context, rawMovieData string) (meta model.MovieMeta, err error)
}
