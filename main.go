package main

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	fmt.Println("Sending message...")

	p := GetProducer()
	topic := "myTopic"
	for _, word := range []string{"Message One", "Message Tow"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	fmt.Println("Getting message...")
	// ...
}
