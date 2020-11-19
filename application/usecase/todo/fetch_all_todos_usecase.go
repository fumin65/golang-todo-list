package todo

import model "todo-api/domain/model/todo"

type FetchAllTodosUseCase struct {
	repository model.TodoRepository
}

func NewFetchAllTodosUseCase(repository model.TodoRepository) *FetchAllTodosUseCase {
	return &FetchAllTodosUseCase{repository: repository}
}

func (uc *FetchAllTodosUseCase) Handle() []*model.Todo {
	return uc.repository.AllTodos()
}
