package commands

import (
	"encoding/json"
	"github.com/DeanThompson/ginpprof"
	"github.com/LiveScraper/app/website-scraping/scraper"
	"github.com/LiveScraper/app/website-scraping/scraper/streaming-services"
	"github.com/LiveScraper/app/website-scraping/scraper/streaming-services/parsers"
	"github.com/LiveScraper/models"
	"github.com/LiveScraper/phttp"
	"github.com/LiveScraper/phttp/client/httpClient"
	"github.com/gin-gonic/gin"
	"log"
	"time"

	"github.com/itsjamie/gin-cors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ServerKey = "server"

var serverConfig *models.Config

var serverCmd = &cobra.Command{
	Use:          "server",
	Short:        "Run the Proj server",
	RunE:         serverCmdF,
	SilenceUsage: true,
}

func init() {
	RootCmd.AddCommand(serverCmd)
	RootCmd.RunE = serverCmdF
}

func initServer() {
	if subConf := viper.Sub(ServerKey); subConf != nil {
		byt, err := json.Marshal(subConf.AllSettings())
		if err != nil {
			log.Panic("unmarshal error")
		}
		if err := json.Unmarshal(byt, &serverConfig); err != nil {
			log.Panic("unable to convert to server conf")
		}
	} else {
		log.Panic("invalid server conf")
	}
}

func serverCmdF(command *cobra.Command, args []string) error {
	gin.SetMode(serverConfig.Mode)
	r := gin.New()
	ginpprof.Wrap(r)

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, PATCH",
		RequestHeaders:  "Origin, Authorization, Content-Type, X-Source",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	networkConnector := httpClient.GetIHttpClient(models.GHttpClient, *models.GUserAgents)

	documentParserFactory := parsers.NewDocumentParserFactory()

	streamingServiceFactory := streaming_services.NewStreamingServiceFactory(documentParserFactory)
	err := streamingServiceFactory.SetStreamingServices(*models.GStreamingServices)
	if err != nil {
		log.Fatal(err.Error())
	}

	scraperService := scraper.NewService(networkConnector, streamingServiceFactory)
	scraper.StreamingServiceHandler(r.Group("movie/"), scraperService)

	phttp.GracefullyServe(r, serverConfig)

	return nil
}
