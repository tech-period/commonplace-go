package api

import (
    "net/http"
)

// HelloHandler は "Hello, World!" を返すハンドラーです。
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, World!"))
}

// HealthCheckHandler はサービスのヘルスチェックを行うハンドラーです。
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}