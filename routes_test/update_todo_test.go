package routes_test

import (
	"bytes"
	"encoding/json"
	"example/ritik/model"
	"example/ritik/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateTodoHandler(t *testing.T) {
	// Create a sample Todo item with ID
	todoID := primitive.NewObjectID()
	sampleTodo := model.Todo{
		ID:        todoID,
		Task:      "Sample Todo",
		Completed: true,
	}
	model.Todos = append(model.Todos, sampleTodo) // Add to in-memory data store

	// Prepare the updated Todo item
	updatedTodo := model.Todo{
		Task:      "Updated Test Todo",
		Completed: true,
	}
	jsonTodo, _ := json.Marshal(updatedTodo)
	req, err := http.NewRequest("PUT", "/todos/"+todoID.Hex(), bytes.NewBuffer(jsonTodo))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	router := gin.Default()
	router.PUT("/todos/:id", routes.UpdateTodoHandler)
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var responseTodo model.Todo
	if err := json.NewDecoder(rr.Body).Decode(&responseTodo); err != nil {
		t.Fatal(err)
	}

	if responseTodo.Task != updatedTodo.Task {
		t.Errorf("handler returned unexpected body: got %v want %v", responseTodo.Task, updatedTodo.Task)
	}
}
