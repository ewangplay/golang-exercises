package main

import (
	"log"
    "time"
    "fmt"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	confirm := make(chan amqp.Confirmation, 1)

	if err := ch.Confirm(false); err != nil {
		log.Println("publisher confirms not supported")
		close(confirm) // confirms not supported, simulate by always nacking
	} else {
		ch.NotifyPublish(confirm)
	}

	messages := make(chan string, 100)

	go func(ch chan string) {
		for i := 0; i < 100; i++ {
			ch <- fmt.Sprintf("Hello %d", i)
		}

		time.Sleep(10 * time.Second)
		close(ch)
	}(messages)

    var msg string
	pending := make(chan string, 1)
	reading := messages

	for {
		select {
		case ack := <-confirm:
			if !ack.Ack {
				log.Printf("Publish message %d: %s failed, retry to publish it...", ack.DeliveryTag, msg)
                pending <- msg
			} else {
				log.Printf("Ack message %d: %s success!", ack.DeliveryTag, msg)
				reading = messages
			}

		case msg = <-pending:

			err = ch.Publish(
				"logs_direct",         // exchange
				"info", // routing key
				false, // mandatory
				false, // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(msg),
				})
			failOnError(err, "Failed to publish a message")

			log.Printf(" [x] Publish %s success!", msg)

		case one_msg, ok := <-reading:
			if !ok {
				log.Printf("channel messages closed\n")
				return
			}
			pending <- one_msg
			reading = nil
		}

	}
}

