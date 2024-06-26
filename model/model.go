package model

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//"go.mongodb.org/mongo-driver/bson/primitive"

/*type Todo struct {
	Todo_Id   string `json:"todo_id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
	Status    string `json:"status"`
}
*/

type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string             `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool               `json:"completed,omitempty" bson:"completed,omitempty"`
}

var Todos = []Todo{
	{ID: primitive.NewObjectID(), Task: "Todo 1", Completed: true},
	{ID: primitive.NewObjectID(), Task: "Todo 2", Completed: false},
}

// GetTodoByID retrieves a todo by ID
func GetTodoByID(id primitive.ObjectID) (Todo, error) {
	for _, todo := range Todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return Todo{}, errors.New("todo not found")
}

// UpdateTodoByID updates a todo by ID
func UpdateTodoByID(id primitive.ObjectID, updatedTodo Todo) error {
	for i, todo := range Todos {
		if todo.ID == id {
			Todos[i] = updatedTodo
			Todos[i].ID = id // Ensure the ID remains unchanged
			return nil
		}
	}
	return errors.New("todo not found")
}
func DeleteTodoByID(id primitive.ObjectID) error {
	for i, todo := range Todos {
		if todo.ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
