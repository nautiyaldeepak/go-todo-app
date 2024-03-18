package router

import (
	"golang-todo-app/healthcheck"
	"golang-todo-app/middleware"
	"golang-todo-app/prometheus"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	// API Endpoint
	router.HandleFunc("/", middleware.HomePage).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todo", middleware.GetAllTodo).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todo", middleware.CreateTodo).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/todo", middleware.UpdateTodo).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", middleware.DeleteTodo).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/alive", healthcheck.Healthcheck).Methods("GET")
	router.Use(prometheus.Metrics)
	router.Path("/metrics").Handler(promhttp.Handler())
	router.PathPrefix("/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	return router
}
