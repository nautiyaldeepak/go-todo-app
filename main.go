package main

import (
	"fmt"
	"golang-todo-app/healthcheck"
	"golang-todo-app/prometheus"
	"golang-todo-app/router"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	hc := healthcheck.Healthcheck()
	fmt.Println(hc)
	r := router.Router()
	// router := mux.NewRouter()
	r.Use(prometheus.PrometheusMiddleware)

	// Prometheus endpoint
	r.Path("/metrics").Handler(promhttp.Handler())
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Starting server on the port 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
