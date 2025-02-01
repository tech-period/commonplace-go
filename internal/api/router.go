package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// ここにエンドポイントとハンドラーを追加します
	// 例: router.HandleFunc("/api/example", ExampleHandler).Methods("GET")

	return router
}
