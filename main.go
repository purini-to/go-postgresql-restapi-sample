package main

import (
	"log"
)

func main() {
	app, cl, err := InitializeApp()
	if err != nil {
		log.Fatal("Initialize failed: %v", err)
	}
	if runErr := app.Run(); runErr != nil {
		log.Fatal("Application run failed: %v", runErr)
		cl()
	}
}
