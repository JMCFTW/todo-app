package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todo "todo-app/domains"
	memory "todo-app/todo/repositorys"
	"todo-app/todo/usecases"
)

type RestController struct {
}

func NewRestController() todo.Controller {
	return &RestController{}
}

func (c *RestController) Execute() {
	repo := memory.NewInMemoryTodoRepository()
	useCase := usecases.NewTodoUseCase(repo)
	router := gin.Default()

	router.POST("api/todos", func(context *gin.Context) {
		var body todo.Todo
		if err := context.ShouldBindJSON(&body); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		useCase.Add(body)
		context.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.GET("api/todos", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"todos": useCase.GetAll()})
	})

	router.PATCH("api/todos/:id/done", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		useCase.DoneById(id)
		context.JSON(http.StatusOK, gin.H{"todos": useCase.GetAll()})
	})

	err := router.Run()
	if err != nil {
		return
	}

}
