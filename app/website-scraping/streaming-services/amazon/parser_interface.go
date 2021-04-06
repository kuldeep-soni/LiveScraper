package amazon

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/enums"
	"github.com/LiveScraper/app/website-scraping/streaming-services/amazon/models"
)

//Every streaming service will have its own interface of parser. All of them might want to fetch different kind of data
type IAmazonDocumentParser interface {
	GetName() enums.ParserType
	TransformRawMovieData(ctx context.Context, rawData string) (meta models.MovieMetaAmazon, err error)
}
