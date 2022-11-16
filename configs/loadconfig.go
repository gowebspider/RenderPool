package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// init Load configs file to config
func init() {
	env := os.Getenv(`env`)
	if env == `` {
		env = `dev`
	}
	viper.SetConfigFile(fmt.Sprintf("configs/%s/config.yaml", env))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		// A fatal error occurred while reading the configuration file
		log.Fatal(err)
	}

	// Listen for configuration changes
	viper.WatchConfig()
}
