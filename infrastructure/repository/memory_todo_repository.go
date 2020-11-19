package repository

import (
	"github.com/google/uuid"
	model "todo-api/domain/model/todo"
)

type MemoryTodoRepository struct {
	todos []*model.Todo
}

func (m *MemoryTodoRepository) AllTodos() []*model.Todo {
	return m.todos
}

func (m *MemoryTodoRepository) TodoOf(id model.TodoId) *model.Todo {
	for _, todo := range m.todos {
		if todo.Id() == id {
			return todo
		}
	}
	return nil
}

func (m *MemoryTodoRepository) Insert(t *model.Todo) {
	m.todos = append(m.todos, t)
}

func (m *MemoryTodoRepository) Save(t *model.Todo) {
	for index, todo := range m.todos {
		if todo.Id() == t.Id() {
			m.todos[index] = t
		}
	}
}

func (m *MemoryTodoRepository) Delete(t *model.Todo) {
	for index, todo := range m.todos {
		if todo.Id() == t.Id() {
			m.todos = append(m.todos[:index], m.todos[index+1:]...)
		}
	}
}

func (m *MemoryTodoRepository) NewId() model.TodoId {
	u, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return model.TodoId(u.String())
}
