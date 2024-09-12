package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func Connect() {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	partitionConsumer, err := consumer.ConsumePartition("task-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to subscribe to partition: %v", err)
	}

	go func() {
		for {
			select {
			case message := <-partitionConsumer.Messages():
				log.Printf("Received message: %s\\n", string(message.Value))
			case err := <-partitionConsumer.Errors():
				log.Printf("Error: %s\\n", err.Err)
			}
		}
	}()

	select {}
}
