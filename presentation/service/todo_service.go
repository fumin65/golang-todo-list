package service

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"todo-api/di/provider"
	"todo-api/domain/model"
)

type TodoServiceImpl struct {
}

func (t *TodoServiceImpl) CreateTodo(c context.Context, req *CreateTodoRequest) (*Todo, error) {

	uc := provider.InitializeCreateTodoUseCase()
	todo := uc.Handle(req.Title)

	createdAt, err := ptypes.TimestampProto(todo.CreateAt())
	if err != nil {
		return nil, err
	}

	res := &Todo{
		Id:        string(todo.Id()),
		Title:     todo.Title(),
		State:     toProtoValue(todo.State()),
		CreatedAt: createdAt,
	}

	return res, nil
}

func (t *TodoServiceImpl) FetchAllTodos(c context.Context, req *FetchAllTodosRequest) (*FetchAllTodosResponse, error) {
	uc := provider.InitializeFetchAllTodosUseCase()
	todos := uc.Handle()

	var resTodos []*Todo
	for _, t := range todos {
		createdAt, err := ptypes.TimestampProto(t.CreateAt())
		if err != nil {
			return nil, err
		}
		resTodos = append(resTodos, &Todo{
			Id:        string(t.Id()),
			Title:     t.Title(),
			State:     toProtoValue(t.State()),
			CreatedAt: createdAt,
		})
	}

	return &FetchAllTodosResponse{
		Todos: resTodos,
	}, nil
}

func toProtoValue(state model.TodoState) Todo_State {
	switch state {
	case model.Completed:
		return Todo_COMPLETED
	case model.NotCompleted:
		return Todo_NOT_COMPLETED
	default:
		return Todo_UNSPECIFIED
	}
}
