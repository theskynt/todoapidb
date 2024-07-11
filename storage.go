package main

import (
	"database/sql"
	"log"
	"strconv"
)

type storage struct {
	conn *sql.DB
}

func NewStorage(conn *sql.DB) *storage {
	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			title TEXT,
			status TEXT
	);
	`

	if _, err := conn.Exec(createTb); err != nil {
		log.Fatal("can't create table", err)
	}

	return &storage{conn}
}

func (s storage) GetAllTodos() ([]Todo, error) {
	rows, err := s.conn.Query("SELECT id, title, status FROM todos")
	if err != nil {
		return []Todo{}, nil
	}

	var todos []Todo
	for rows.Next() {
		var id int
		var title, status string
		err := rows.Scan(&id, &title, &status)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}

		todos = append(todos, Todo{
			ID:     id,
			Title:  title,
			Status: status,
		})
	}

	return todos, nil
}

func (s storage) GetTodoById(rowId string) (*Todo, error) {
	q := "SELECT id, title, status FROM todos WHERE id=$1"
	row := s.conn.QueryRow(q, rowId)

	var id int
	var title, status string
	err := row.Scan(&id, &title, &status)
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:     id,
		Title:  title,
		Status: status,
	}, nil
}

func (s storage) PostTodo(title, status string) (*Todo, error) {
	q := "INSERT INTO todos (title, status) values ($1, $2)  RETURNING id"
	row := s.conn.QueryRow(q, title, status)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	return s.GetTodoById(strconv.Itoa(id))
}

func (s storage) PutTodo(rowId, title, status string) (*Todo, error) {
	q := "UPDATE todos SET title=$2, status=$3 WHERE id=$1;"
	if _, err := s.conn.Exec(q, rowId, title, status); err != nil {
		return nil, err
	}

	return s.GetTodoById(rowId)
}

func (s storage) DeleteTodo(rowId string) string {
	q := "DELETE FROM todos WHERE id=$1;"
	if _, err := s.conn.Exec(q, rowId); err != nil {
		return "fail"
	}

	return "succes"
}

func (s storage) PutStatusTodo(rowId, status string) (*Todo, error) {
	q := "UPDATE todos SET status=$2 WHERE id=$1;"
	if _, err := s.conn.Exec(q, rowId, status); err != nil {
		return nil, err
	}

	return s.GetTodoById(rowId)
}

func (s storage) PutTitleTodo(rowId, title string) (*Todo, error) {
	q := "UPDATE todos SET title=$2 WHERE id=$1;"
	if _, err := s.conn.Exec(q, rowId, title); err != nil {
		return nil, err
	}

	return s.GetTodoById(rowId)
}
