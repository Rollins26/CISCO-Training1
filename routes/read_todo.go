package routes

import (
	"context"
	getcollection "example/ritik/collections"
	database "example/ritik/databases"
	"example/ritik/model"

	//"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadOnetodo(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var todoCollection = getcollection.GetCollection(DB, "todos")

	todoId := c.Param("todo_id")
	var result model.Todo

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(todoId)

	err := todoCollection.FindOne(ctx, bson.M{"todo_id": objId}).Decode(&result)

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success!", "Data": res})

}
func ReadTodoHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Assuming there's some logic to retrieve the todo item from the database
	todo, err := model.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}
