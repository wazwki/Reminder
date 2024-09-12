package main

import (
	"log"
	"log/slog"
	"net/http"
	"taskservice/internal/handlers"
	"taskservice/internal/storage"
	"taskservice/pkg/kafka"
	"taskservice/pkg/logging"
)

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

	if err := http.ListenAndServe(host, mux); err != nil {
		log.Fatal(err)
	}
}
