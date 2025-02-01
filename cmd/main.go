package main

import (
    "log"
    "net/http"
    "commonplace-go/internal/api"
)

func main() {
    router := api.NewRouter()
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}