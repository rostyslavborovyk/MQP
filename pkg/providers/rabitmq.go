package providers

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"reflect"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func createConnection(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func createChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

type RabbitMQProvider struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQProvider(url string) *RabbitMQProvider {
	conn := createConnection(url)
	ch := createChannel(conn)
	return &RabbitMQProvider{
		conn: conn,
		ch:   ch,
	}
}

func (p *RabbitMQProvider) Close() error {
	err := p.ch.Close()
	if err != nil {
		return err
	}
	err = p.conn.Close()
	if err != nil {
		return err
	}
	return nil
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

func (p *RabbitMQProvider) PushMessage(queueName string, messageType string, message interface{}) bool {
	var body []byte
	if reflect.ValueOf(message).Kind() == reflect.String {
		body = []byte(message.(string))
	} else {
		var err error
		if body, err = json.Marshal(message); err != nil {
			failOnError(err, "Failed to marshal message to json")
		}
	}
	if err := p.ch.PublishWithContext(
		context.Background(),
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: messageType,
			Body:        body,
		}); err != nil {
		failOnError(err, "Failed to push message")
	}
	return true
}

func (p RabbitMQProvider) GetConsumerChanel(queueName string) (<-chan amqp.Delivery, error) {
	consumer, err := p.ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
