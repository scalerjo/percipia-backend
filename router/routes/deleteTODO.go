package routes

import (
	"backend/database"
	_ "fmt"

	"github.com/gin-gonic/gin"
)

// Route to create a new TODO item.
// Inserts into database
func DeleteTODO(c *gin.Context) {
	id := c.Param("id")

	_, err := database.DB.Exec("DELETE FROM todo WHERE id=$1", id);
	if err != nil {
		c.Status(400)
		return
	}

	c.Status(200)
}
