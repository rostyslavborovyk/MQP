package providers

type Provider interface {
	Close() error
	PushMessage(queueName string, messageType string, message interface{}) bool
	CreateQueue(name string) bool
}
