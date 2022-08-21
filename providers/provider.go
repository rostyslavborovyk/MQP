package providers

type Provider interface {
	Init()
	Close()
	PushMessage(queueName string, message string) bool
	CreateQueue(name string) bool
}
