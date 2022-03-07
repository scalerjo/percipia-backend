package router


import (
	"github.com/gin-gonic/gin"

	"backend/router/routes"
	"backend/database"
)

func Start() {
	database.ConnectDB()

	router := gin.Default()

	// Bind requested resources to internal routes
	router.GET("/api/todo", routes.GetAllTODO)
	router.POST("/api/todo", routes.CreateTODO)
	router.PUT("/api/todo", routes.UpdateTODO)
	router.DELETE("/api/todo/:id", routes.DeleteTODO)

	// Start the router on port 8888
	router.Run("0.0.0.0:8888")
}
