package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvConfig struct {
	mysqlUri string
}

func getEnvConfig() EnvConfig {
	godotenv.Load()

	config := EnvConfig{}
	config.mysqlUri = os.Getenv("MYSQL_URI")
	if config.mysqlUri == "" {
		log.Fatal("Missing env variables")
	}

	return config
}
