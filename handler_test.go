package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func setupRouter() *gin.Engine {
	db := setupTestDB()
	store := NewStorage(db)
	h := NewHandler(store)

	r := gin.Default()
	r.GET("/todos", h.GetTodos)
	r.GET("/todos/:id", h.GetTodo)
	r.POST("/todos", h.CreateTodo)
	r.PUT("/todos/:id", h.UpdateTodo)
	r.DELETE("/todos/:id", h.DeleteTodo)
	r.PATCH("/todos/:id/status", h.UpdateStatusTodo)
	r.PATCH("/todos/:id/title", h.UpdateTitleTodo)

	return r
}
func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}