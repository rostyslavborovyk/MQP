// consumer is a script intended to test a mqp producer
package main

import (
	"fmt"
	"github.com/rostyslavborovyk/MQP/pkg/providers"
	"log"
)

func main() {
	provider := providers.NewRabbitMQProvider("amqp://guest:guest@localhost:5672/")

	delivery, err := provider.Consume("random-queue")
	if err != nil {
		log.Panicf("error when consuming message from queue %v", err)
	}
	fmt.Printf("%v\n", delivery)
}
