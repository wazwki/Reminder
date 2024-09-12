package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

var Producer sarama.AsyncProducer

func Connect() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	var err error

	Producer, err = sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}

	go func() {
		for {
			select {
			case err := <-Producer.Errors():
				log.Printf("Failed to send message: %v", err)
			case success := <-Producer.Successes():
				log.Printf("Message sent to partition %d with offset %d\\n", success.Partition, success.Offset)
			}
		}
	}()

	select {}
}
