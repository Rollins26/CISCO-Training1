package routes

import (
	"context"
	"encoding/json"
	getcollection "example/ritik/collections"
	database "example/ritik/databases"
	"example/ritik/model"
	"log"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func Createtodo(c *gin.Context) {
	var DB = database.ConnectDB()
	var todoCollection = getcollection.GetCollection(DB, "todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	todo := new(model.Todo)
	defer cancel()

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	postPayload := model.Todo{
		ID:        todo.ID,
		Task:      todo.Task,
		Completed: todo.Completed,
	}

	result, err := todoCollection.InsertOne(ctx, postPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Assuming there's some logic to create the todo item and store it
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
