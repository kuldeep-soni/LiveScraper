package scraper_test

import (
	"context"
	"github.com/LiveScraper/app/website-scraping/enums"
	"github.com/LiveScraper/app/website-scraping/service"
	streaming_services "github.com/LiveScraper/app/website-scraping/streaming-services"
	"github.com/LiveScraper/global"
	"github.com/LiveScraper/phttp/client/httpClient"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestScraperSvc(t *testing.T){
	networkConnector := httpClient.GetIHttpClient("mock", []string{"Mozilla/5.0"})
	streamingServiceFactory := streaming_services.NewStreamingServiceFactory()

	ss := global.StreamingServices{{
		Name: "amazon",
		URL: "http://www.amazon.de/gp/product",
		ParserType: "amazon-parser-1",
	}}
	err := streamingServiceFactory.SetStreamingServices(ss)
	if err != nil {
		log.Fatal(err.Error())
	}

	scraperService := service.NewScraper(networkConnector, streamingServiceFactory)

	meta, err := scraperService.GetMovieMeta(context.Background(), enums.Amazon, "B00FCM7N9C")
	assert.NoError(t, err)
	assert.Equal(t, "Der Exorzist", meta.Title)
}
