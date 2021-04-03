package commands

import (
	"encoding/json"
	"github.com/DeanThompson/ginpprof"
	"github.com/LiveScraper/models"
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

	//root context this would be passed to each and every call
	//ctx := context.Background()

	//logger := plog.NewLogger(&plog.LoggerConfiguration{
	//	EnableConsole: true,
	//	ConsoleJson:   true,
	//})

	//Standard logger for this whole service
	//plog.RedirectStdLog(logger)

	gin.SetMode(serverConfig.Mode)
	r := gin.New()
	ginpprof.Wrap(r)

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	//r.Use(middleware.RequestLogger(logger.StdLog(plog.String("source", "request_log"))),
	//	middleware.XRequestID(), middleware.HandleError())
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, PATCH",
		RequestHeaders:  "Origin, Authorization, Content-Type, X-Source",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	models.GracefullyServe(r, serverConfig)

	return nil
}