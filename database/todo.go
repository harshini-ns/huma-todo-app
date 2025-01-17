package database

import (
	"database/sql"
	"fmt"
	"time"
)

type Todo struct {
	ID      int64  `json:"id,omitempty"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type todoDAO struct {
	*sql.DB
}

var TodoDAO *todoDAO

func InitTodoDAO() {
	TodoDAO = &todoDAO{
		DB: pgClient,
	}
}

func (h *todoDAO) AddTodo(todo *Todo) (*Todo, error) {
	todo.ID = time.Now().UnixNano() / int64(time.Millisecond)
	query := `INSERT INTO todo (id, title, content) VALUES ($1, $2, $3)`
	fmt.Print("inserted todo to DB : ", todo)
	_, err := h.DB.Exec(query, todo.ID, todo.Title, todo.Content)
	if err != nil {
		fmt.Printf("Error inserting into database: %v\n", err)
		return nil, err
	}
	fmt.Println("Todo successfully added")
	return todo, nil
}

func (h *todoDAO) GetTodo(id int64) (*Todo, error) {
	query := `SELECT id, title, content FROM todo WHERE id = $1`
	row := h.DB.QueryRow(query, id)
	var todo Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &todo, nil

}

func (h *todoDAO) DeleteTodo(id int64) (*Todo, error) {
	todo, err := h.GetTodo(id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, nil
	}
	query := `DELETE FROM todo WHERE id = $1`
	_, err = h.DB.Exec(query, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("Todo successfully deleted")
	return todo, nil
}

func (h *todoDAO) UpdateTodo(id int64, title string, content string) (*Todo, error) {
	todo, err := h.GetTodo(id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, nil
	}
	query := `UPDATE todo SET title = $2, content = $3 WHERE id = $1`
	_, err = h.DB.Exec(query, id, title, content)
	if err != nil {
		return nil, err
	}
	todo.Title = title
	todo.Content = content
	return todo, nil
}
