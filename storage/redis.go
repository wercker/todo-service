package storage

import (
	"encoding/json"

	"gopkg.in/redis.v3"

	"github.com/mies/todo-service/todo"
)

type RedisStateStore struct {
	client *redis.Client
}

func NewRedisStateStore(client *redis.Client) *RedisStateStore {
	return &RedisStateStore{client}
}

func (r RedisStateStore) Insert(t *todo.Todo) error {
	j, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = r.client.LPush("todos", string(j)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisStateStore) GetTodos() ([]*todo.Todo, error) {
	result, _ := r.client.LRange("todos", -100, 100).Result()
	list := make([]*todo.Todo, len(result))
	for key, value := range result {
		var t todo.Todo
		err := json.Unmarshal([]byte(value), &t)
		if err != nil {
			return nil, err
		}
		//t := todo.Todo{Name: value}
		list[key] = &t
	}
	return list, nil
}
