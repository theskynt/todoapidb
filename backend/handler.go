package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	st *storage
}

func NewHandler(st *storage) *handler {
	return &handler{st}
}

func (h handler) GetTodos(c *gin.Context) {
	getTodos, err := h.st.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getTodos)
}

func (h handler) GetTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "id is required")
		return
	}

	getTodo, err := h.st.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, getTodo)
}

func (h handler) CreateTodo(c *gin.Context) {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	newTodo, err := h.st.PostTodo(todo.Title, todo.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}

func (h handler) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "id is required")
		return
	}

	var todo Todo
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	newTodo, err := h.st.PutTodo(id, todo.Title, todo.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}

func (h handler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "id is required")
		return
	}

	deleteTodo := h.st.DeleteTodo(id)
	c.JSON(http.StatusCreated, deleteTodo)
}

func (h handler) UpdateStatusTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "id is required")
		return
	}

	var todo Todo
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	newTodo, err := h.st.PutStatusTodo(id, todo.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}

func (h handler) UpdateTitleTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "id is required")
		return
	}

	var todo Todo
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	newTodo, err := h.st.PutTitleTodo(id, todo.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}
