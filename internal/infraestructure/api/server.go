package api

import (
	"golang-clean-arch-reference/internal/infraestructure/api/routes"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	routes.HealthRoutes(r)
	routes.CustomerRoutes(r)

	r.Run()
}
