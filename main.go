package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	todos := []Todo{}
	router := gin.Default()

	router.POST("api/todos", func(context *gin.Context) {
		var body Todo
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

	err := router.Run()
	if err != nil {
		return
	}
}
