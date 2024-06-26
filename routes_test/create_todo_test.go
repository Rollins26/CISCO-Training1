package routes_test

import (
	"bytes"
	"encoding/json"
	"example/ritik/model"
	"example/ritik/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTodoHandler(t *testing.T) {
	// Create a new HTTP request
	newTodo := model.Todo{
		Task:      "Test Todo",
		Completed: true,
	}
	jsonTodo, _ := json.Marshal(newTodo)
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonTodo))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.CreateTodoHandler)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body
	var createdTodo model.Todo
	if err := json.NewDecoder(rr.Body).Decode(&createdTodo); err != nil {
		t.Fatal(err)
	}

	if createdTodo.Task != newTodo.Task {
		t.Errorf("handler returned unexpected body: got %v want %v", createdTodo.Task, newTodo.Task)
	}
}
