package main

import (
	"github.com/gin-gonic/gin"
	"school/todo"
)

func main() {

	todoHandler := todo.TodoHandler{}
	r := gin.Default()
	r.POST("/api/todos",todoHandler.PostTodosHandler)
	r.GET("/api/todos", todoHandler.GetListTodosHandler)
	r.GET("/api/todos/:id", todoHandler.GetTodosByIdHandler)
	r.PUT("/api/todos/:id", todoHandler.PutUpdateTodoHandler)
	r.DELETE("/api/todos/:id", todoHandler.DeleteTodosHandler)
	r.Run(":1234")
}