package provider

import (
	"todo-api/domain/model"
	"todo-api/infrastructure/repository"
)

var repo = &repository.MemoryTodoRepository{}

func NewTodoRepository() model.TodoRepository {
	return repo
}
