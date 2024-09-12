package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"taskservice/internal/storage"
	"taskservice/internal/task_types"
	"taskservice/pkg/kafka"

	"github.com/IBM/sarama"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	var newtask task_types.Task
	json.NewDecoder(r.Body).Decode(&newtask)

	if storage.DB == nil {
		log.Fatal("Database connection is not initialized")
	}

	_, err := storage.DB.Exec(`UPDATE tasks SET title=$1, description=$2 WHERE id=$3`, newtask.Title, newtask.Description, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: "task-topic",
		Value: sarama.StringEncoder("task_updated"),
	}

	kafka.Producer.Input() <- msg

	w.WriteHeader(http.StatusOK)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newtask task_types.Task
	json.NewDecoder(r.Body).Decode(&newtask)

	var lastInsertId int

	if storage.DB == nil {
		log.Fatal("Database connection is not initialized")
	}

	err := storage.DB.QueryRow("INSERT INTO tasks(title, description) VALUES($1, $2) RETURNING id", newtask.Title, newtask.Description).Scan(&lastInsertId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(lastInsertId)

	msg := &sarama.ProducerMessage{
		Topic: "task-topic",
		Value: sarama.StringEncoder("task_created"),
	}

	kafka.Producer.Input() <- msg

	w.WriteHeader(http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	if storage.DB == nil {
		log.Fatal("Database connection is not initialized")
	}

	_, err := storage.DB.Exec(`DELETE FROM tasks WHERE id=$1`, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: "task-topic",
		Value: sarama.StringEncoder("task_deleted"),
	}

	kafka.Producer.Input() <- msg

	w.WriteHeader(http.StatusOK)
}
