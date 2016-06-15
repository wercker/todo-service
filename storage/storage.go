package storage

import (
	todo "github.com/mies/todo-service/todo"
	//"errors"
)

type StateStore interface {
	GetTodos() ([]*todo.Todo, error)
	Insert(*todo.Todo) error
}
