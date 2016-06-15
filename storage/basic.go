package storage

import (
	"sync"

	todo "github.com/mies/todo-service/todo"
)

type MemStore struct {
	db   []*todo.Todo
	lock *sync.Mutex
}

func NewMemStore() *MemStore {
	db := []*todo.Todo{}
	return &MemStore{
		db,
		&sync.Mutex{},
	}
}

func (m *MemStore) GetTodos() ([]*todo.Todo, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.db, nil
}

func (m *MemStore) Insert(t *todo.Todo) error {
	m.lock.Lock()
	m.db = append(m.db, t)
	m.lock.Unlock()
	return nil
}
