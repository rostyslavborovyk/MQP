package populator

import (
	config2 "github.com/rostyslavborovyk/MQP/config"
	"github.com/rostyslavborovyk/MQP/pkg/providers"
	"log"
)

type Controller struct {
	populators []*Populator
}

func (c *Controller) Init() {
	config := config2.GetConfig()
	for _, service := range config.Services {
		switch service.Type {
		case "rabbitmq":
			provider := &providers.RabbitMQProvider{}
			provider.Init(service.Url)
			c.populators = append(c.populators, &Populator{
				provider:     provider,
				queuesConfig: service.Queues,
			})
		default:
			log.Panicf("Unable to handle service type %s", service.Type)
		}
	}
}

func (c Controller) Start() {
	for _, p := range c.populators {
		p.Run()
	}
}

func (c *Controller) Close() {
	for _, p := range c.populators {
		p.Close()
	}
}
