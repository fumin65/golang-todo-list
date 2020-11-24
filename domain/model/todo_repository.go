package model

type TodoRepository interface {
	AllTodos() []*Todo

	TodoOf(id TodoId) *Todo

	Insert(t *Todo)

	Save(t *Todo)

	Delete(t *Todo)

	NewId() TodoId
}
