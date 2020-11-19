package todo

import "time"

type State int

const (
	Completed    State = iota
	NotCompleted State = iota
)

type Todo struct {
	id        TodoId
	title     string
	state     State
	createdAt time.Time
}

func (t *Todo) Id() TodoId {
	return t.id
}

func NewTodo(id TodoId, title string) *Todo {
	return &Todo{id: id, title: title, state: NotCompleted, createdAt: time.Now()}
}

func (t *Todo) CreateAt() time.Time {
	return t.createdAt
}

func (t *Todo) State() State {
	return t.state
}

func (t *Todo) Title() string {
	return t.title
}

func (t *Todo) SetTitle(title string) {
	t.title = title
}

func (t *Todo) Complete() {
	if t.state == NotCompleted {
		t.state = Completed
	}
}

func (t *Todo) Uncomplete() {
	if t.state == Completed {
		t.state = NotCompleted
	}
}
