package main

import (
	"log"
)

func main() {
	app, cl, err := InitializeApp()
	if err != nil {
		log.Fatal("Initialize failed: %v", err)
	}
	if runErr := app.Start(); runErr != nil {
		log.Fatal("Application start failed: %v", runErr)
		cl()
	}
	cl()
}
