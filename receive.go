package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"

	proto_msg "queue/proto-msg"
)

// receive cli command settings
var receiveCmd = &cobra.Command{
	Use:   "receive [binding key]...",
	Short: "Receive RabbitMQ message",
	Args:  cobra.MinimumNArgs(1),
	Run:   runReceive,
}

func init() {
	rootCmd.AddCommand(receiveCmd)
}

func runReceive(cmd *cobra.Command, args []string) {
	routes := args
	var amqpUrl string
	if username != "" && password != "" {
		amqpUrl = fmt.Sprintf("%s://%s:%s@%s:%d/", scheme, username, password, host, port)
	} else {
		amqpUrl = fmt.Sprintf("%s://%s:%d/", scheme, host, port)
	}
	config := amqp.Config{
		Vhost: vhost,
	}
	conn, err := amqp.DialConfig(amqpUrl, config)
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
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	for _, route := range routes {
		log.Printf("Binding queue %s to exchange %s with routing key %q", q.Name, exchange, route)
		err = ch.QueueBind(
			q.Name,   // queue name
			route,    // routing key
			exchange, // exchange
			false,
			nil)
		if err != nil {
			log.Fatalf("Failed to bind a queue: %s", err)
		}
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for delivery := range msgs {
			topic := delivery.RoutingKey
			log.Printf("topic: %v", topic)
			msg := delivery.Body
			data := &proto_msg.UserMsg{}
			err := proto.Unmarshal(msg, data)
			if err != nil {
				log.Printf("Unmarshal MatchMakerWildcowMsg error: %v", err)
			} else {
				switch m := data.Msg.(type) {
				case *proto_msg.UserMsg_UserBasic:
					log.Printf("UserMsg_UserBasic: %v", m)
				case *proto_msg.UserMsg_UserBalance:
					log.Printf("UserMsg_UserBalance: %v", m)
				case *proto_msg.UserMsg_UserStartGameInfo:
					log.Printf("UserMsg_UserStartGameInfo: %v", m)
				case *proto_msg.UserMsg_UserToken:
					log.Printf("UserMsg_UserToken: %v", m)
				default:
					log.Printf("Unknown: %v", m)
				}
			}
		}
	}()

	log.Printf("Waiting for logs. To exit press CTRL+C")
	<-forever
}
