//+build wireinject

package provider

import (
	"github.com/google/wire"
	"todo-api/application/usecase/todo"
)

func InitializeCreateTodoUseCase() *todo.CreateTodoUseCase {
	wire.Build(todo.NewCreateTodoUseCase, NewTodoRepository)
	return &todo.CreateTodoUseCase{}
}

func InitializeFetchAllTodosUseCase() *todo.FetchAllTodosUseCase {
	wire.Build(todo.NewFetchAllTodosUseCase, NewTodoRepository)
	return &todo.FetchAllTodosUseCase{}
}
