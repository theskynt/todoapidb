package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	s := NewStorage(db)
	h := NewHandler(s)

	r := gin.Default()
	r.LoadHTMLGlob("./*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	todoRoute := r.Group("/api/v1/todos")
	todoRoute.GET(":id", h.GetTodo)
	todoRoute.GET("", h.GetTodos)
	todoRoute.POST("", h.CreateTodo)
	todoRoute.PUT(":id", h.UpdateTodo)
	todoRoute.DELETE(":id", h.DeleteTodo)
	todoRoute.PATCH(":id/actions/status", h.UpdateStatusTodo)
	todoRoute.PATCH(":id/actions/title", h.UpdateTitleTodo)

	srv := http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	idleConnsClosed := make(chan struct{})

	go func() {
		<-ctx.Done()
		fmt.Println("Shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			slog.Error(err.Error())
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
		}
	}

	<-idleConnsClosed
	fmt.Println("Bye!!!")
}
