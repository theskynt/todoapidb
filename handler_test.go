package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestHandlerDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
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

var db *sql.DB

func setupRouter() *gin.Engine {
	db = setupTestHandlerDB()
	store := NewStorage(db)
	h := NewHandler(store)

	r := gin.Default()
	r.GET("/api/v1/todos", h.GetTodos)
	r.GET("/api/v1/todos/:id", h.GetTodo)
	r.POST("/api/v1/todos", h.CreateTodo)
	r.PUT("/api/v1/todos/:id", h.UpdateTodo)
	r.DELETE("/api/v1/todos/:id", h.DeleteTodo)
	r.PATCH("/api/v1/todos/:id/actions/status", h.UpdateStatusTodo)
	r.PATCH("/api/v1/todos/:id/actions/title", h.UpdateTitleTodo)

	return r
}

func TestHandlerGetTodos(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/api/v1/todos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandlerGetTodo(t *testing.T) {
	router := setupRouter()

	store := NewStorage(db)
	todo, _ := store.PostTodo("Test title", "Test status")
	fmt.Println(todo.ID)

	req, _ := http.NewRequest("GET", "/api/v1/todos/"+strconv.Itoa(todo.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandlerCreateTodo(t *testing.T) {
	router := setupRouter()

	todo := Todo{Title: "Test title", Status: "Test status"}
	jsonValue, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/api/v1/todos", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHandlerUpdateTodo(t *testing.T) {
	router := setupRouter()

	store := NewStorage(db)
	todo, _ := store.PostTodo("Test title", "Test status")

	updatedTodo := Todo{Title: "Updated title", Status: "Updated status"}
	jsonValue, _ := json.Marshal(updatedTodo)
	req, _ := http.NewRequest("PUT", "/api/v1/todos/"+strconv.Itoa(todo.ID), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHandlerDeleteTodo(t *testing.T) {
	router := setupRouter()

	store := NewStorage(db)
	todo, _ := store.PostTodo("Test title", "Test status")

	req, _ := http.NewRequest("DELETE", "/api/v1/todos/"+strconv.Itoa(todo.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHandlerUpdateStatusTodo(t *testing.T) {
	router := setupRouter()

	store := NewStorage(db)
	todo, _ := store.PostTodo("Test title", "Test status")

	updatedStatus := Todo{Status: "New status"}
	jsonValue, _ := json.Marshal(updatedStatus)
	req, _ := http.NewRequest("PATCH", "/api/v1/todos/"+strconv.Itoa(todo.ID)+"/actions/status", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHandlerUpdateTitleTodo(t *testing.T) {
	router := setupRouter()

	store := NewStorage(db)
	todo, _ := store.PostTodo("Test title", "Test status")

	updatedTitle := Todo{Title: "New title"}
	jsonValue, _ := json.Marshal(updatedTitle)
	req, _ := http.NewRequest("PATCH", "/api/v1/todos/"+strconv.Itoa(todo.ID)+"/actions/title", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
