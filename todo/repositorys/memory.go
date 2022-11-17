package memory

import (
	todo "todo-app/domains"
)

type InMemoryToDoRepository struct {
	todos []todo.Todo
}

func (r *InMemoryToDoRepository) Add(todo todo.Todo) (int, error) {
	todo.ID = len(r.todos) + 1
	r.todos = append(r.todos, todo)
	return todo.ID, nil
}

func (r *InMemoryToDoRepository) DoneById(id int) (int, error) {
	r.todos[id-1].Done = true
	return id, nil
}

func (r *InMemoryToDoRepository) GetAll() []todo.Todo {
	return r.todos
}

func NewInMemoryTodoRepository() todo.TodoRepository {
	var todos []todo.Todo
	return &InMemoryToDoRepository{todos}
}
