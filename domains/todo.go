package todo

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

type TodoRepository interface {
	Add(todo Todo) (int, error)
	DoneById(id int) (int, error)
	GetAll() []Todo
}

type TodoUsecase interface {
	Add(todo Todo) (int, error)
	DoneById(id int) (int, error)
	GetAll() []Todo
}
