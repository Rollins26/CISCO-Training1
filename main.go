package main

import (

	//"net/http"

	"example/ritik/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/todos", routes.Createtodo)

	// called as localhost:3000/getOne/{id}
	router.GET("/todos/:todo_id", routes.ReadOnetodo)

	// called as localhost:3000/update/{id}
	router.PUT("/todos/:todo_id", routes.Updatetodo)

	// called as localhost:3000/delete/{id}
	router.DELETE("/todos/:todo_id", routes.Deletetodo)

	router.Run("localhost: 8081")
}
