package main

import (
	"github.com/gin-gonic/gin"
	rest "todo-app/todo/controllers"
	memory "todo-app/todo/repositorys"
	"todo-app/todo/usecases"
)

func main() {
	repo := memory.NewInMemoryTodoRepository()
	useCase := usecases.NewTodoUseCase(repo)
	router := gin.Default()
	rest.NewRestController(router, useCase)
	err := router.Run()
	if err != nil {
		return
	}
}
