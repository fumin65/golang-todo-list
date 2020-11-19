package provider

import (
	"todo-api/domain/model/todo"
	"todo-api/infrastructure/repository"
)

var repo = &repository.MemoryTodoRepository{}

func NewTodoRepository() todo.TodoRepository {
	return repo
}
