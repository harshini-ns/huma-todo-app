package handlers

import (
	"todo/database"

	"github.com/danielgtaylor/huma/v2"
)

func addTodo(todo *database.Todo) (*database.Todo, error) {
	if todo.Title == "" || todo.Content == "" {
		return nil, huma.Error400BadRequest("Title and content cannot be empty")
	}
	return database.TodoDAO.AddTodo(todo)
}

func updateTodo(todo *database.Todo) (*database.Todo, error) {
	if todo.ID == 0 || todo.Title == "" || todo.Content == "" {
		return nil, huma.Error400BadRequest(" invalid todo")
	}
	todoDB, err := database.TodoDAO.GetTodo(todo.ID)
	if err != nil || todoDB == nil {
		return nil, huma.Error400BadRequest(" todo not found")
	}
	respTodo, err := database.TodoDAO.UpdateTodo(todo.ID, todo.Title, todo.Content)
	if err != nil {
		return nil, huma.Error400BadRequest("failed")
	}
	return respTodo, nil
}
