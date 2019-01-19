package main

import (
	"log"
)

func main() {
	server, cl, err := InitializeServer()
	if err != nil {
		log.Fatal("Initialize failed", err)
	}
	if runErr := server.Start(); runErr != nil {
		cl()
		log.Fatal("Server start failed", runErr)
	}
	cl()
}
