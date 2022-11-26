package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todo "todo-app/domains"
)

type RestController struct {
	todoUsecase todo.TodoUsecase
}

func NewRestController(engine *gin.Engine, usecase todo.TodoUsecase) {
	controller := &RestController{
		todoUsecase: usecase,
	}
	engine.POST("api/todos", controller.PostTodo)

	engine.GET("api/todos", controller.GetTodo)

	engine.PATCH("api/todos/:id/done", controller.PatchTodo)
}

func (controller *RestController) PostTodo(context *gin.Context) {
	var body todo.Todo
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.todoUsecase.Add(body)
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *RestController) GetTodo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"todos": controller.todoUsecase.GetAll()})
}

func (controller *RestController) PatchTodo(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	controller.todoUsecase.DoneById(id)
	context.JSON(http.StatusOK, gin.H{"todos": controller.todoUsecase.GetAll()})
}
