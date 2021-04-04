package commands

import (
	"encoding/json"
	"github.com/LiveScraper/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

type Command = cobra.Command

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "website-scraping",
	Short: "Service to scrape web pages",
	Long:  `This service is responsible for scraping data from websites and transforming them to useful data`,
}

func init() {
	cobra.OnInitialize(initConfig, initServer)
	RootCmd.PersistentFlags().StringP("config", "c", "app/website-scraping/cmd/config/configuration.json", "Configuration file to use.")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
}

func initConfig() {
	cfgFile := viper.Get("config").(string)
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Panic("Can't read config, ", err)
	}

	streamingServicesBytes, err := json.Marshal(viper.Get("streaming-services"))
	if err != nil{
		log.Panic("Can't read streaming-services key values, ", err)
	}
	pStreamingServices := &models.StreamingServices{}
	err = json.Unmarshal(streamingServicesBytes, pStreamingServices)
	if err != nil{
		log.Panic("Can't unmarshal streaming service data to struct ", err)
	}

	models.GStreamingServices = pStreamingServices

	models.GHttpClient = viper.GetString("http_client")

	userAgentsBytes, err := json.Marshal(viper.Get("user-agents"))
	if err != nil{
		log.Panic("Can't read user-agents key values, ", err)
	}
	pUserAgents := &models.UserAgents{}
	err = json.Unmarshal(userAgentsBytes, pUserAgents)
	if err != nil{
		log.Panic("Can't unmarshal streaming service data to struct ", err)
	}

	models.GUserAgents = pUserAgents
}
