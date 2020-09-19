package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://rabbit:inthehole@192.168.0.9:5672")
	if err != nil {
	    log.Fatal(err)
	}
	defer conn.Close()

	mq, err := conn.Channel()
	if err != nil {
	    log.Fatal(err)
	}
	defer mq.Close()

	queue, err := mq.QueueDeclare("test", false, false, false, false, nil)
	if err != nil {
	    log.Fatal(err)
	}

	for i := 0;; i++ {
		err := mq.Publish("", queue.Name, false, false, amqp.Publishing{
			ContentType:     "text/plain",
			Body:            []byte("message " + strconv.Itoa(i)),
		})
		if err != nil {
		    fmt.Println(err)
		}
		time.Sleep(time.Second)
	}
}