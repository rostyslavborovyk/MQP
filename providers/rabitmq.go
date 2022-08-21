package providers

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func CreateConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

type RabbitMQProvider struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (p *RabbitMQProvider) Init() {
	p.conn = CreateConnection()
	p.ch = CreateChannel(p.conn)
}

func (p *RabbitMQProvider) Close() {
	p.ch.Close()
	p.conn.Close()
}

func (p *RabbitMQProvider) CreateQueue(name string) bool {
	_, err := p.ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue")
	return true
}

func (p RabbitMQProvider) PushMessage(queueName string, message string) bool {
	if err := p.ch.PublishWithContext(
		context.Background(),
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		}); err != nil {
		failOnError(err, "Failed to push message")
	}
	return true
}
