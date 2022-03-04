package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Route to create a new TODO item.
// Inserts into database
func CreateTODO(c *gin.Context) {
	c.String(http.StatusOK, "Hello I am working")
}
