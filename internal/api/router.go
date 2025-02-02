package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// ルートの設定
	// 例: router.GET("/api/example", ExampleHandler)
	router.GET("/api/example", HealthCheckHandler)

	return router
}
