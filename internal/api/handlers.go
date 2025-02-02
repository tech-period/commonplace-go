package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
