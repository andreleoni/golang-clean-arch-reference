package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthRoutes(routes *gin.Engine) {
	routes.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
