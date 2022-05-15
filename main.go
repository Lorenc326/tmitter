package main

import (
	"github.com/Lorenc326/tmitter/db"
	"log"
)

func main() {
	log.Println("starting server")
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
}
