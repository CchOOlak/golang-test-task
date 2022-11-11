package messaging

import (
	"encoding/json"
	"os"

	"github.com/streadway/amqp"
)

type RabbitClient struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

var Client *RabbitClient

const MessageQueue = "messages"

func Init() {
	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	Client.Connection = connectRabbitMQ

	// Create channel
	channelRabbitMQ, err := Client.Connection.Channel()
	if err != nil {
		panic(err)
	}
	Client.Channel = channelRabbitMQ

	// Declare Queue
	_, err = Client.Channel.QueueDeclare(
		MessageQueue, // queue name
		true,         // durable
		false,        // auto delete
		false,        // exclusive
		false,        // no wait
		nil,          // arguments
	)
	if err != nil {
		panic(err)
	}
}

func (c *RabbitClient) Close() {
	c.Channel.Close()
	c.Connection.Close()
}

func (c *RabbitClient) Produce(msg interface{}) error {
	// Create a message to publish.
	byteMsg, _ := json.Marshal(msg)
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        byteMsg,
	}

	// Attempt to publish a message to the queue.
	if err := c.Channel.Publish(
		"",           // exchange
		MessageQueue, // queue name
		false,        // mandatory
		false,        // immediate
		message,      // message to publish
	); err != nil {
		return err
	}
	return nil
}

func (c *RabbitClient) Consume(processor func([]byte)) {
	messages, err := c.Channel.Consume(
		MessageQueue, // queue name
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // arguments
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			go processor(message.Body)
		}
	}()

	<-forever
}
