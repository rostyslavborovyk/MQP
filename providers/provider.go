package providers

type Provider interface {
	Init(url string)
	Close()
	PushMessage(queueName string, message string) bool
	CreateQueue(name string) bool
}
