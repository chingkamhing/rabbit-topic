package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

// send cli command settings
var sendCmd = &cobra.Command{
	Use:   "send [message]",
	Short: "Send RabbitMQ message",
	Args:  cobra.ExactArgs(1),
	Run:   runSend,
}

func init() {
	rootCmd.AddCommand(sendCmd)
}

func runSend(cmd *cobra.Command, args []string) {
	body := args[0]
	var amqpUrl string
	if username != "" && password != "" {
		amqpUrl = fmt.Sprintf("%s://%s:%s@%s:%d/", scheme, username, password, host, port)
	} else {
		amqpUrl = fmt.Sprintf("%s://%s:%d/", scheme, host, port)
	}
	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}

	err = ch.Publish(
		"logs_topic", // exchange
		route,        // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}

	log.Printf("Sent %s", body)
}
