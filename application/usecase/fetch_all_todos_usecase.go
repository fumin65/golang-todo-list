package usecase

import (
	"todo-api/domain/model"
)

type FetchAllTodosUseCase struct {
	repository model.TodoRepository
}

func NewFetchAllTodosUseCase(repository model.TodoRepository) *FetchAllTodosUseCase {
	return &FetchAllTodosUseCase{repository: repository}
}

func (uc *FetchAllTodosUseCase) Handle() []*model.Todo {
	return uc.repository.AllTodos()
}
