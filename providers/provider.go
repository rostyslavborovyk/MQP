package providers

type Provider interface {
	Init(url string)
	Close()
	PushMessage(queueName string, messageType string, message interface{}) bool
	CreateQueue(name string) bool
}
