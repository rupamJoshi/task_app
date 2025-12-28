package router

import (
	"log"
	"net/http"
	"os"
	"task_app/handler/task"
	task_repo "task_app/repo/task"
	task_service "task_app/service/task"

	"github.com/julienschmidt/httprouter"
)

func NewServer() *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	return &http.Server{Addr: "localhost:" + port, Handler: newHandler()}
}

func addRoutes(r *httprouter.Router) {

	taskRepo := task_repo.NewTaskRepo()
	taskService := task_service.NewService(taskRepo)
	handler := task.NewHandler(taskService)
	r.POST("/task", handler.Create)
	r.GET("/task/:id", handler.Get)
	r.GET("/tasks", handler.GetAll)
	r.PUT("/task/:id", handler.Update)
	r.DELETE("/task/:id", handler.Delete)
}

func newHandler() http.Handler {
	r := httprouter.New()
	addRoutes(r)

	return r
}
