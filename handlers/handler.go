package handlers

import (
	"context"
	"fmt"
	"net/http"
	"todo/database"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

func addTodo(todo *database.Todo) (*database.Todo, error) {
	return database.TodoDAO.AddTodo(todo)
}

type inputTodo struct {
	Body *database.Todo
}

type outputTodo struct {
	Body *database.Todo
}

func InitHandler() {
	router := chi.NewMux()

	api := humachi.New(router, huma.DefaultConfig("Todo API", "1.0.0"))

	huma.Register(api, huma.Operation{
		OperationID: "create-todo",
		Method:      http.MethodPost,
		Path:        "/todo",
		Summary:     "Create a new todo item",
	}, func(ctx context.Context, input *inputTodo) (*outputTodo, error) {
		todo := input.Body
		resp := &outputTodo{}
		fmt.Printf("Received Todo Payload: %+v\n", todo)
		if todo.Title == "" || todo.Content == "" {
			return nil, huma.Error400BadRequest("Title and content cannot be empty")
		}
		respTodo, err := addTodo(todo)
		if err != nil {
			return nil, huma.Error400BadRequest("error in addTodo")
		}
		resp.Body = respTodo
		return resp, nil
	})

	//
	huma.Register(api, huma.Operation{
		OperationID: "get-todo",
		Method:      http.MethodGet,
		Path:        "/todo/{id}",
		Summary:     "to get todo by id",
	}, func(ctx context.Context, input *struct {
		ID int64 `path:"id" doc:"unique id to get"`
	}) (*outputTodo, error) {
		resp := &outputTodo{}
		todo, err := database.TodoDAO.GetTodo(input.ID)
		if err != nil {
			return nil, huma.Error400BadRequest(" todo not found")
		}
		resp.Body = todo
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "delete-todo",
		Method:      http.MethodDelete,
		Path:        "/todo/{id}",
		Summary:     "to delete a todo by id",
	}, func(ctx context.Context, input *struct {
		ID int64 `path:"id" doc:"id to delete"`
	}) (*outputTodo, error) {
		resp := &outputTodo{}
		todo, err := database.TodoDAO.DeleteTodo(input.ID)
		if err != nil {
			return nil, huma.Error400BadRequest(" todo not found")
		}
		resp.Body = todo
		return resp, nil
	})

	//update todo
	huma.Register(api, huma.Operation{
		OperationID: "update a todo",
		Method:      http.MethodPut,
		Path:        "/todo",
		Summary:     "to update a todo by id",
	}, func(ctx context.Context, input *inputTodo) (*outputTodo, error) {
		todo := input.Body
		if todo.Title == "" {
			return nil, huma.Error400BadRequest(" invalid todo")
		}
		respTodo, err := database.TodoDAO.UpdateTodo(todo.ID, todo.Title, todo.Content)
		if err != nil {
			return nil, huma.Error400BadRequest("failed")
		}
		return &outputTodo{
			Body: respTodo,
		}, nil
	})

	http.ListenAndServe("127.0.0.1:8888", router)
}
