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
	router.GET("/todo", routes.GetAllTODO)
	router.POST("/todo", routes.CreateTODO)
	router.PUT("/todo", routes.UpdateTODO)
	router.DELETE("/todo/:id", routes.DeleteTODO)

	// Start the router on port 8888
	router.Run("0.0.0.0:8888")
}
