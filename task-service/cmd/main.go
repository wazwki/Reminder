package main

import (
	"log"
	"log/slog"
	"net/http"
	"path/filepath"
	"strings"
	"taskservice/internal/handlers"
	"taskservice/internal/storage"
	"taskservice/pkg/kafka"
	"taskservice/pkg/logging"
	"taskservice/pkg/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @title Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /

const (
	host = "localhost:8080"
)

func main() {
	logging.LogInit()
	slog.SetDefault(logging.Logger)

	err := storage.Connect()
	if err != nil {
		log.Fatal(err)
	}
	Postgres := storage.DB
	_, err = Postgres.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	go kafka.Connect()
	log.Println("Started")
	mux := http.NewServeMux()

	mux.HandleFunc("PUT /tasks/{id}", handlers.UpdateTask)
	mux.HandleFunc("POST /tasks", handlers.CreateTask)
	mux.HandleFunc("DELETE /tasks/{id}", handlers.DeleteTask)
	mux.HandleFunc("/swagger/", serveSwagger)
	mux.Handle("/metrics", promhttp.Handler())

	metricsMux := metrics.MetricsMiddleware(mux)

	if err := http.ListenAndServe(host, metricsMux); err != nil {
		log.Fatal(err)
	}
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	if p == "" {
		p = "index.html"
	}
	p = filepath.Join("cmd", "swagger", "docs", p)
	http.ServeFile(w, r, p)
}
