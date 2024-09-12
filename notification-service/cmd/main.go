package main

import (
	"log"
	"notificationservice/pkg/kafka"
)

func main() {
	go kafka.Connect()
	log.Println("Started")
	select {}
}
