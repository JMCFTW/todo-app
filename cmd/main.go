package main

import (
	rest "todo-app/todo/controllers"
	memory "todo-app/todo/repositorys"
	"todo-app/todo/usecases"
)

func main() {
	repo := memory.NewInMemoryTodoRepository()
	useCase := usecases.NewTodoUseCase(repo)
	err := rest.NewRestController(useCase)
	if err != nil {
		return
	}
}
