package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://rabbit:inthehole@192.168.0.9:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mq, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer mq.Close()

	queue, err := mq.QueueDeclare(
		"test", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := mq.Consume(
		queue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for d := range msgs {
			fmt.Printf("received message: %s\n", d.Body)
		}
	}()

	forever := make(chan bool)
	<-forever
}