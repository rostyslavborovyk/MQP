package providers

//go:generate mockgen -source=provider.go -destination=mocks/mock.go

type Provider interface {
	Close() error
	PushMessage(queueName string, messageType string, message interface{}) bool
	CreateQueue(name string) bool
}
