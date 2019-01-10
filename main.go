package main

import (
	"log"

	"github.com/purini-to/go-postgresql-restapi-sample/server"
)

func main() {
	app, cl, err := server.InitializeApp()
	if err != nil {
		log.Fatal("Initialize failed: %v", err)
	}
	if runErr := app.Run(); runErr != nil {
		log.Fatal("Application run failed: %v", runErr)
		cl()
	}
}
