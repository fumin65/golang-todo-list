package todo

import model "todo-api/domain/model/todo"

type CreateTodoUseCase struct {
	repository model.TodoRepository
}

func NewCreateTodoUseCase(repository model.TodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{repository: repository}
}

func (uc *CreateTodoUseCase) Handle(title string) *model.Todo {
	id := uc.repository.NewId()
	todo := model.NewTodo(id, title)
	uc.repository.Insert(todo)
	return todo
}
