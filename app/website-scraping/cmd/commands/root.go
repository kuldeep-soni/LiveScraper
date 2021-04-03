package commands

import (
	"encoding/json"
	"fmt"
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
	Short: "Service to scrape we pages",
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

	data, err := json.Marshal(viper.Get("streaming-services"))
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(data)
}
