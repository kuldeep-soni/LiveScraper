package scraper_test

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/scraper"
	streaming_services "github.com/LiveScraper/app/website-scraping/scraper/streaming-services"
	"github.com/LiveScraper/app/website-scraping/scraper/streaming-services/parsers"
	"github.com/LiveScraper/models"
	"github.com/LiveScraper/phttp/client/httpClient"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestScraperSvc(t *testing.T){
	networkConnector := httpClient.GetIHttpClient("mock")

	documentParserFactory := parsers.NewDocumentParserFactory()

	streamingServiceFactory := streaming_services.NewStreamingServiceFactory(documentParserFactory)

	ss := models.StreamingServices{{
		Name: "amazon",
		URL: "http://www.amazon.de/gp/product",
		ParserType: "amazon-parser-1",
	}}
	err := streamingServiceFactory.SetStreamingServices(ss)
	if err != nil {
		log.Fatal(err.Error())
	}

	scraperService := scraper.NewService(networkConnector, streamingServiceFactory)

	for i:=0; i<10 ; i++{
		_, err := scraperService.GetMovieMeta(context.Background(), streaming_services.Amazon, "B00FCM7N9C")
		assert.NoError(t, err)
	}
}
