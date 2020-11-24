//+build wireinject

package provider

import (
	"github.com/google/wire"
	"todo-api/application/usecase"
)

func InitializeCreateTodoUseCase() *usecase.CreateTodoUseCase {
	wire.Build(usecase.NewCreateTodoUseCase, NewTodoRepository)
	return &usecase.CreateTodoUseCase{}
}

func InitializeFetchAllTodosUseCase() *usecase.FetchAllTodosUseCase {
	wire.Build(usecase.NewFetchAllTodosUseCase, NewTodoRepository)
	return &usecase.FetchAllTodosUseCase{}
}
