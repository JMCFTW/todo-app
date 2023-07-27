package usecases

import (
	"github.com/stretchr/testify/suite"
	"testing"
	todo "todo-app/domains"
	memory "todo-app/todo/repositorys"
)

func TestToDoUseCasesTestSuite(t *testing.T) {
	suite.Run(t, new(ToDoUseCasesTestSuite))
}

type ToDoUseCasesTestSuite struct {
	suite.Suite
	todousecase todo.TodoUsecase
}

func (s *ToDoUseCasesTestSuite) SetupTest() {
	repo := memory.NewInMemoryTodoRepository()
	s.todousecase = NewTodoUseCase(repo)
}

func (s *ToDoUseCasesTestSuite) Test_get_todo() {
	s.Equal([]todo.Todo([]todo.Todo(nil)), s.todousecase.GetAll())
}

func (s *ToDoUseCasesTestSuite) Test_add_todo() {
	expectedId, _ := s.todousecase.Add(todo.Todo{ID: 1, Title: "do something", Done: false, Body: "some description"})
	s.Equal(1, expectedId)
}

func (s *ToDoUseCasesTestSuite) Test_done_todo() {
	s.todousecase.Add(todo.Todo{ID: 1, Title: "do something", Done: false, Body: "some description"})
	expectedId, _ := s.todousecase.DoneById(1)
	s.Equal(1, expectedId)
}
