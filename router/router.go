package router


import (
	"github.com/gin-gonic/gin"

	"backend/router/routes"
)

func Start() {
	router := gin.Default()

	// Bind requested resources to internal routes
	router.GET("/", routes.CreateTODO)

	// Start the router on port 8888
	router.Run("0.0.0.0:8888")
}
