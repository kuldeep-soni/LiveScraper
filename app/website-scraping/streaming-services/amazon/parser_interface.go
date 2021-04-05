package amazon

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/enums"
	"github.com/LiveScraper/app/website-scraping/streaming-services/amazon/models"
)

type IAmazonDocumentParser interface {
	GetName() enums.ParserType
	TransformRawMovieData(ctx context.Context, rawData string) (meta models.MovieMetaAmazon, err error)
}
