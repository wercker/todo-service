package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"gopkg.in/urfave/cli.v1"

	storage "github.com/mies/todo-service/storage"
	todo "github.com/mies/todo-service/todo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

var (
	store storage.StateStore
	err   error
)

var (
	startServer = cli.Command{
		Name:  "start",
		Usage: "start the todo service",
		Action: func(c *cli.Context) {
			opts := ParseOpts(c)
			store, err = getStateStore(c, opts)
			if err != nil {
				fmt.Printf("error %v\n", err)
				panic(err)
			}
			l, err := net.Listen("tcp", fmt.Sprintf(":%s", opts.ServicePort))
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer()
			todo.RegisterDoSomethingServer(s, &server{})
			s.Serve(l)
		},
	}
)

func (s *server) AddTodo(ctx context.Context, t *todo.Todo) (*todo.Empty, error) {
	err := store.Insert(t)
	if err != nil {
		return &todo.Empty{}, err
	}
	return &todo.Empty{}, nil
}

func (s *server) ListTodos(ctx context.Context, t *todo.Empty) (*todo.TodoList, error) {
	result, err := store.GetTodos()
	if err != nil {
		return nil, err
	}
	list := &todo.TodoList{
		Todos: result,
	}
	return list, nil
}

func main() {

	app := cli.NewApp()
	app.Name = "todo-service"
	app.Usage = "service for adding and listing todos"
	app.Commands = []cli.Command{
		startServer,
	}
	app.Flags = GlobalFlags()
	app.Run(os.Args)
}
