package main

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"todo-api/presentation/service"
)

func main() {
	port, err := net.Listen("tcp", ":19003")

	if err != nil {
		println(err)
		os.Exit(0)
	}

	server := grpc.NewServer()
	service.RegisterTodoServiceServer(server, &service.TodoServiceImpl{})

	err = server.Serve(port)
	println(err)
}
