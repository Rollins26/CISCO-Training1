package routes_test

import (
	"encoding/json"
	"example/ritik/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteTodoHandler(t *testing.T) {
	// Create a new Gorilla mux router
	r := mux.NewRouter()

	// Register the DeleteTodoHandler with the router
	r.HandleFunc("/todos/{id}", routes.DeleteTodoHandler).Methods("DELETE")

	// Generate a primitive.ObjectID for testing
	objID := primitive.NewObjectID()

	// Create a request with DELETE method for /todos/{id}
	req, err := http.NewRequest("DELETE", "/todos/"+objID.Hex(), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Serve the HTTP request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check if the todo is deleted (this requires your actual DeleteTodoHandler logic)
	var responseMessage map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&responseMessage); err != nil {
		t.Fatal(err)
	}

	expectedMessage := "Todo deleted"
	if message, ok := responseMessage["message"]; !ok || message != expectedMessage {
		t.Errorf("handler returned unexpected body: got %v want %v", message, expectedMessage)
	}
}
