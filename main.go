package main

import (
	"fmt"
	"github.com/Lorenc326/tmitter/db"
	"github.com/Lorenc326/tmitter/tweet"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	log.Println("initializing server")
	config := getEnvConfig()

	log.Println("connecting to the db")
	sqlClient, err := db.Connect(config.mysqlUri)
	if err != nil {
		log.Fatal(err)
	}
	db.Client = sqlClient
	defer db.Client.Close()

	// remove from start up in case app has more than 1 instances
	if err = db.MigrateUp(db.Client); err != nil {
		log.Fatal(err)
	}

	server := buildServer()
	address := fmt.Sprintf(":%s", config.appPort)
	server.Logger.Fatal(server.Start(address))
}

func buildServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tweets/:id", tweet.GetTweet)
	e.GET("/tweets", tweet.GetTweetsByUser)

	return e
}
