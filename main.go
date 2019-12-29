package main

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	fmt.Println("Sending message...")

	p := GetProducer()
	topic := "myTopic"
	for _, word := range []string{"MessageOne", "MessageTow"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	fmt.Println("Getting message...")

	c := GetConsumer()
	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
