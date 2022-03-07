package routes

import (
	"github.com/gin-gonic/gin"
	"backend/database"
	"backend/types"
	_ "fmt"
)

// Route to create a new TODO item.
// Inserts into database
func CreateTODO(c *gin.Context) {
	var newTODO types.TODO

    // Call BindJSON to bind the received JSON
    if err := c.BindJSON(&newTODO); err != nil {
		c.Status(401)
        return
    }

	if (len(newTODO.Text) == 0) || (len(newTODO.Text) > 199) || (newTODO.Time == 0) {
		c.Status(402)
        return
	}

	_, err := database.DB.Exec("INSERT INTO todo(text, time) VALUES($1, $2)", newTODO.Text, newTODO.Time);
	if err != nil {
		c.Status(400)
		return
	}

	c.Status(200)
}
