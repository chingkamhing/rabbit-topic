package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"

	proto_msg "queue/proto-msg"
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
	name := args[0]
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

	msg := &proto_msg.UserMsg{}
	if len(name) < 10 {
		msg.Msg = &proto_msg.UserMsg_UserBasic{
			UserBasic: &proto_msg.UserBasic{
				UserID:        101,
				GameID:        101,
				Accounts:      "Account " + name,
				NickName:      name,
				FaceID:        "",
				Gender:        "",
				Nullity:       false,
				LastLogonDate: 0,
			},
		}
	} else {
		msg.Msg = &proto_msg.UserMsg_UserBalance{
			UserBalance: &proto_msg.UserBalance{
				UserID:            1,
				Score:             2,
				Revenue:           3,
				InsureScore:       4,
				AllWinScore:       5,
				BetTotal:          6,
				BetCount:          7,
				EffectiveBetTotal: 8,
				JackpotScore:      9,
				TmScore:           10,
				JackpotBet:        11,
			},
		}
	}
	body, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalf("JoinSuccess Marshal error: %v", err)
	}
	err = ch.Publish(
		exchange, // exchange
		route,    // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}

	log.Printf("Sent %s", body)
}
