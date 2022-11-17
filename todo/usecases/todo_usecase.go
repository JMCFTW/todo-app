package usecases

import todo "todo-app/domains"

type TodoUseCase struct {
	todoRepo todo.TodoRepository
}

func NewTodoUseCase(todoRepo todo.TodoRepository) todo.TodoUsecase {
	return &TodoUseCase{todoRepo}
}

func (t *TodoUseCase) Add(todo todo.Todo) (int, error) {
	id, err := t.todoRepo.Add(todo)
	return id, err
}

func (t *TodoUseCase) DoneById(id int) (int, error) {
	id, err := t.todoRepo.DoneById(id)
	return id, err
}

func (t *TodoUseCase) GetAll() []todo.Todo {
	return t.todoRepo.GetAll()
}
