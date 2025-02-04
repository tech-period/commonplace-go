package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

// @Summary ヘルスチェックエンドポイント
// @Description システムの健全性を確認するエンドポイント
// @Tags health
// @Accept json
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /example [get]
func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
