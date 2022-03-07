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

	// var response []JsonResponse
	var todos []types.TODO

	// Foreach movie
	for rows.Next() {
		var id int
		var text string
		var time int64

		err = rows.Scan(&id, &text, &time)

		if err != nil {
			c.Status(400)
			return
		}
		todo := types.TODO{
			ID: id,
			Text: text,
			Time: time,
		}
		todos = append(todos, todo)
	}
	c.IndentedJSON(200, todos)
}
