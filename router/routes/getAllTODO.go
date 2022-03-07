package routes

import (
	"backend/database"
	"backend/types"
	"github.com/gin-gonic/gin"
)

// Route to create a new TODO item.
// Inserts into database
func GetAllTODO(c *gin.Context) {

	rows, err := database.DB.Query("SELECT * FROM todo")
	if err != nil {
		c.Status(400)
		return
	}

	// result variable
	var todos []types.TODO

	// loop through each row and construct TODO objects
	for rows.Next() {
		var id int
		var text string
		var time int64

		// scan the row
		err = rows.Scan(&id, &text, &time)

		if err != nil {
			c.Status(400)
			return
		}

		// construct the TODO object
		todo := types.TODO{
			ID: id,
			Text: text,
			Time: time,
		}

		// append to result variable
		todos = append(todos, todo)
	}

	// Send result back to client
	c.IndentedJSON(200, todos)
}
