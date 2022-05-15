package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvConfig struct {
	mysqlUri string
	appPort  string
}

func getEnvConfig() EnvConfig {
	godotenv.Load()

	config := EnvConfig{}
	config.mysqlUri = mustGetEnv("MYSQL_URI")
	config.appPort = mustGetEnv("APP_PORT")

	return config
}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("missing env variable %s", key)
	}
	return val
}
