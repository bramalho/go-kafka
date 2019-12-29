package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var producer *kafka.Producer
var consumer *kafka.Consumer

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	initProducer()
	initConsumer()
}

func initProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST") + ":9092",
	})
	if err != nil {
		log.Fatal(err)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	producer = p
}

func initConsumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST"),
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	consumer = c
}

// GetProducer instance
func GetProducer() *kafka.Producer {
	return producer
}

// GetConsumer instance
func GetConsumer() *kafka.Consumer {
	return consumer
}
