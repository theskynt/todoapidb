package main

import (
	"database/sql"
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

type TodoTest struct {
	ID     int
	Title  string
	Status string
}

func setupTestDB() *sql.DB {
	db, err := sql.Open("sqlite", "file:TestCreateSkillHandlerIT?mode=memory&cache=shared")
	if err != nil {
		log.Fatal(err)
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		status TEXT
	);
	`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	return db
}

func TestGetAllTodos(t *testing.T) {
	db := setupTestDB()
	store := NewStorage(db)

	store.PostTodo("Test title 1", "Test status 1")
	store.PostTodo("Test title 2", "Test status 2")

	todos, err := store.GetAllTodos()
	if err != nil {
		t.Fatalf("Failed to get todos: %v", err)
	}

	assert.Equal(t, 2, len(todos))
	assert.Equal(t, "Test title 1", todos[0].Title)
	assert.Equal(t, "Test status 1", todos[0].Status)
}

func TestGetTodoById(t *testing.T) {
	db := setupTestDB()
	store := NewStorage(db)

	// เพิ่มข้อมูลตัวอย่าง
	todo, err := store.PostTodo("Test title", "Test status")
	if err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	todoFetched, err := store.GetTodoById(strconv.Itoa(todo.ID))
	if err != nil {
		t.Fatalf("Failed to get todo: %v", err)
	}

	assert.Equal(t, todo.ID, todoFetched.ID)
	assert.Equal(t, todo.Title, todoFetched.Title)
	assert.Equal(t, todo.Status, todoFetched.Status)
}

func TestPostTodo(t *testing.T) {
	db := setupTestDB()
	store := NewStorage(db)

	todo, err := store.PostTodo("Test title", "Test status")
	if err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	assert.NotNil(t, todo)
	assert.Equal(t, "Test title", todo.Title)
	assert.Equal(t, "Test status", todo.Status)
}

func TestPutTodo(t *testing.T) {
	db := setupTestDB()
	store := NewStorage(db)

	// เพิ่มข้อมูลตัวอย่าง
	todo, err := store.PostTodo("Test title", "Test status")
	if err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	updatedTodo, err := store.PutTodo(strconv.Itoa(todo.ID), "Updated title", "Updated status")
	if err != nil {
		t.Fatalf("Failed to update todo: %v", err)
	}

	assert.Equal(t, todo.ID, updatedTodo.ID)
	assert.Equal(t, "Updated title", updatedTodo.Title)
	assert.Equal(t, "Updated status", updatedTodo.Status)
}

func TestDeleteTodo(t *testing.T) {
	db := setupTestDB()
	store := NewStorage(db)

	todo, err := store.PostTodo("Test title", "Test status")
	if err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	result := store.DeleteTodo(strconv.Itoa(todo.ID))
	assert.Equal(t, "succes", result)

	deletedTodo, err := store.GetTodoById(strconv.Itoa(todo.ID))
	assert.Nil(t, deletedTodo)
	assert.NotNil(t, err)
}

func TestPutStatusTodo(t *testing.T) {
	db := setupTestDB()
	store := NewStorage(db)

	todo, err := store.PostTodo("Test title", "Test status")
	if err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	updatedTodo, err := store.PutStatusTodo(strconv.Itoa(todo.ID), "New status")
	if err != nil {
		t.Fatalf("Failed to update status: %v", err)
	}

	assert.Equal(t, "New status", updatedTodo.Status)
}

func TestPutTitleTodo(t *testing.T) {
	db := setupTestDB()
	store := NewStorage(db)

	todo, err := store.PostTodo("Test title", "Test status")
	if err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	updatedTodo, err := store.PutTitleTodo(strconv.Itoa(todo.ID), "New title")
	if err != nil {
		t.Fatalf("Failed to update title: %v", err)
	}

	assert.Equal(t, "New title", updatedTodo.Title)
}
