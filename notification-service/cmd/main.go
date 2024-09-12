package main

import (
	"log"
	"log/slog"
	"notificationservice/pkg/kafka"
	"notificationservice/pkg/logging"
)

func main() {
	logging.LogInit()
	slog.SetDefault(logging.Logger)

	go kafka.Connect()
	log.Println("Started")

	select {}
}
