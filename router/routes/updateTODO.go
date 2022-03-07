package routes

import (
	"github.com/gin-gonic/gin"
	"backend/database"
	"backend/types"
	_ "fmt"
)

// Route to create a new TODO item.
// Inserts into database
func UpdateTODO(c *gin.Context) {
	var newTODO types.TODO

    // Call BindJSON to bind the received JSON
    if err := c.BindJSON(&newTODO); err != nil {
		c.Status(401)
        return
    }

	// Input Validation
	if (len(newTODO.Text) == 0) || (newTODO.Time == 0) || (newTODO.ID == 0) {
		c.Status(402)
        return
	}

	// query the database and check for errors
	_, err := database.DB.Exec("UPDATE todo SET text=$2, time=$3 WHERE id=$1", newTODO.ID, newTODO.Text, newTODO.Time);
	if err != nil {
		c.Status(400)
		return
	}

	// Signal to the client that the opperation was successful
	c.Status(200)
}
