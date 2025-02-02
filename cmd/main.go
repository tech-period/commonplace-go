package main

import (
	"commonplace-go/internal/api"
	"log"
)

func main() {
	router := api.NewRouter()
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
