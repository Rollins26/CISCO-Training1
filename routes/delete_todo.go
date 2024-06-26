package routes

import (

	//model "CRUD_API/CRUD_API/model"

	"context"
	"encoding/json"
	getcollection "example/ritik/collections"
	database "example/ritik/databases"
	"example/ritik/model"

	//"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Deletetodo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	todoId := c.Param("todo_id")

	var todoCollection = getcollection.GetCollection(DB, "todos")
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(todoId)
	result, err := todoCollection.DeleteOne(ctx, bson.M{"todo_id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully", "Data": res})

}

// DeleteTodoHandler handles the deletion of a Todo item

// DeleteTodoHandler handles the deletion of a Todo item
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the todo ID from the URL path
	vars := mux.Vars(r)
	id := vars["id"]

	// Convert id string to primitive.ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Assuming there's some logic to delete the todo item from the database
	err = model.DeleteTodoByID(objID)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Send a success response
	response := map[string]string{"message": "Todo deleted"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
