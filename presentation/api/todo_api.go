package api

import (
	"github.com/labstack/echo"
	"net/http"
	"todo-api/di/provider"
	model "todo-api/domain/model/todo"
	"todo-api/presentation/api/req"
	"todo-api/presentation/api/res"
)

func CreateTodo(c echo.Context) error {
	newTodo := new(req.NewTodo)

	if err := c.Bind(newTodo); err != nil {
		return err
	}

	uc := provider.InitializeCreateTodoUseCase()
	todo := uc.Handle(newTodo.Title)

	json := res.Todo{
		Id:        string(todo.Id()),
		Title:     todo.Title(),
		State:     toTextValue(todo.State()),
		CreatedAt: todo.CreateAt(),
	}
	return c.JSON(http.StatusOK, json)
}

func FetchAllTodo(c echo.Context) error {
	uc := provider.InitializeFetchAllTodosUseCase()
	todos := uc.Handle()

	var resTodos []res.Todo
	for _, t := range todos {
		resTodos = append(resTodos, res.Todo{
			Id:        string(t.Id()),
			Title:     t.Title(),
			State:     toTextValue(t.State()),
			CreatedAt: t.CreateAt(),
		})
	}

	return c.JSON(http.StatusOK, resTodos)
}

func toTextValue(state model.State) string {
	switch state {
	case model.Completed:
		return "完了"
	case model.NotCompleted:
		return "未完了"
	default:
		return "不明"
	}
}
