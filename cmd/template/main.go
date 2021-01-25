package main

import (
	"encoding/json"
	"github.com/MaxPolarfox/http_server/pkg/helpers/environment"
	"github.com/MaxPolarfox/http_server/pkg/helpers/logger"
	"github.com/MaxPolarfox/http_server/pkg/template"
	"io/ioutil"
	"os"
)

const ServiceName = "template"

func main() {
	env := environment.GetEnvironment()
	options := loadEnvironmentConfig(env)

	logger := logger.NewLogger(os.Stdout, options.Logger)

	service := template.NewService(options, env, logger)
	service.Start()
}

// loadEnvironmentConfig will use the environment string and concatenate to a proper config file to use
func loadEnvironmentConfig(env string) template.Options {
	configFile := "config/" + ServiceName + "/" + env + ".json"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		panic(err)
	}
	return parseConfigFile(configFile)
}

func parseConfigFile(configFile string) template.Options {
	var opts template.Options
	byts, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byts, &opts)
	if err != nil {
		panic(err)
	}

	return opts
}
