package main

// @title Commonplace Go API
// @version 1.0
// @description This is a sample server for Commonplace Go API.
// @host localhost:8080
// @BasePath /api

import (
	_ "commonplace-go/docs" // Swaggerドキュメントをインポート
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
