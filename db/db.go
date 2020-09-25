package db

import (
	"fmt"

	"github.com/doniacld/simple-web-api/todo"
)

var (
	currentID int
	TodosList todo.Todos
)

// Give us some seed data
func init() {
	RepoCreateTodo(todo.Todo{Name: "Doing more programming"})
	RepoCreateTodo(todo.Todo{Name: "Doing more yoga"})
}

// RepoCreateTodo creates a todo
func RepoCreateTodo(t todo.Todo) todo.Todo {
	currentID += 1
	t.ID = currentID
	TodosList = append(TodosList, t)
	return t
}

// RepoFindTodo finds a todo
func RepoFindTodo(id int) (todo.Todo, error) {
	if id > len(TodosList) {
		return todo.Todo{}, fmt.Errorf("could not find todo with id %d to delete", id)
	}
	for _, t := range TodosList {
		if t.ID == id {
			return t, nil
		}
	}
	return todo.Todo{}, fmt.Errorf("could not find todo with id %d", id)
}

// RepoDestroyTodo deletes a todo
func RepoDestroyTodo(id int) error {
	for i, t := range TodosList {
		if t.ID == id {
			TodosList = append(TodosList[:i], TodosList[i+1:]...)
			return nil
		}

	}
	return fmt.Errorf("could not find todo with id %d to delete", id)
}

// RepoRetrieveTodos returns the list of todos
func RepoRetrieveTodos() todo.Todos {
	return TodosList
}
