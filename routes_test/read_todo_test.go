package routes_test

import (
	"encoding/json"
	"example/ritik/model"
	"example/ritik/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestReadTodoHandler(t *testing.T) {
	// Create a sample Todo item with ID
	todoID := primitive.NewObjectID()
	sampleTodo := model.Todo{
		ID:        todoID,
		Task:      "Sample Todo",
		Completed: true,
	}
	model.Todos = append(model.Todos, sampleTodo) // Add to in-memory data store

	req, err := http.NewRequest("GET", "/todos/"+todoID.Hex(), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	router := gin.Default()
	router.GET("/todos/:id", routes.ReadTodoHandler)
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var readTodo model.Todo
	if err := json.NewDecoder(rr.Body).Decode(&readTodo); err != nil {
		t.Fatal(err)
	}

	if readTodo.ID != todoID {
		t.Errorf("handler returned unexpected body: got %v want %v", readTodo.ID, todoID)
	}
}
