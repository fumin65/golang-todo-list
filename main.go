package main

import (
	"github.com/labstack/echo"
	"todo-api/presentation/api"
)

func main() {
	e := echo.New()
	e.GET("/todos", api.FetchAllTodo)
	e.POST("/todos", api.CreateTodo)
	e.Logger.Fatal(e.Start(":8080"))
}
