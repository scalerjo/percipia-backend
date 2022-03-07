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

	// Query the database and check for errors
	_, err := database.DB.Exec("DELETE FROM todo WHERE id=$1", id);
	if err != nil {
		c.Status(400)
		return
	}

	// Signal to the client that the opperation was successful
	c.Status(200)
}
