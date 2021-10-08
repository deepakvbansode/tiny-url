package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "WISHER"

type WebServerConfig struct {
	Port           string `required:"true" split_words:"true" default:"50051"`
	RoutePrefix    string `required:"false" split_words:"true" default:"/"`
}

type Config struct {
	WebServerConfig *WebServerConfig
	//Mongo                mongo.DatabaseHelper
}

//FromEnv ...
func FromEnv() (cfg *WebServerConfig, err error) {
	fromFileToEnv()

	cfg = &WebServerConfig{}

	err = envconfig.Process(envPrefix, cfg)

	if err != nil {

		return nil, nil
	}

	return cfg, nil
}

func fromFileToEnv() {

	cfgFilename := "../../etc/config/config.localhost.env"

	err := godotenv.Load(cfgFilename)
	if err != nil {
		fmt.Println("No config files found to load to env. Defaulting to environment.", err)
	}

}
