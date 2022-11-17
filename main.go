package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-app/domains"
)

func main() {
	var todos []todo.Todo
	router := gin.Default()

	router.POST("api/todos", func(context *gin.Context) {
		var body todo.Todo
		if err := context.ShouldBindJSON(&body); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		body.ID = len(todos) + 1
		todos = append(todos, body)
		context.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.GET("api/todos", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"todos": todos})
	})

	router.PATCH("api/todos/:id/done", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		todos[id-1].Done = true
		context.JSON(http.StatusOK, gin.H{"todos": todos})
	})

	err := router.Run()
	if err != nil {
		return
	}
}
