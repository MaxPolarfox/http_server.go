package environment

import (
	"fmt"
	"os"
)

const EnvironmentVariable = "APP_ENV"

var environments = map[string]string{
	"production":        "production",
	"development":       "development",
	"development-local": "development-local",
}

// GetEnvironment returns the system EnvironmentVariable (APP_ENV) environment variable
func GetEnvironment() string {
	appEnv := os.Getenv(EnvironmentVariable)
	if appEnv == "" {
		panic(fmt.Errorf("APP_ENV environment variable is not set"))
	}

	environment, ok := environments[os.Getenv(EnvironmentVariable)]
	if ok {
		return environment
	} else {
		panic(fmt.Errorf("environment %s not supported", environment))
	}
}