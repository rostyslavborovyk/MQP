package populator

import (
	"fmt"
	config2 "github.com/rostyslavborovyk/MQP/config"
	"github.com/rostyslavborovyk/MQP/pkg/providers"
	"log"
)

type Controller struct {
	populators []*Populator
}

func NewController() *Controller {
	controller := Controller{}
	config := config2.GetConfig()
	for _, service := range config.Services {
		switch service.Type {
		case "rabbitmq":
			provider := providers.NewRabbitMQProvider(service.Url)
			controller.populators = append(controller.populators, &Populator{
				provider:     provider,
				queuesConfig: service.Queues,
			})
		default:
			log.Panicf("Unable to handle service type %s", service.Type)
		}
	}
	return &controller
}

func (c Controller) Start() {
	for _, p := range c.populators {
		p.Run()
	}
	fmt.Println("Started MQP!")
}

func (c *Controller) Close() {
	for _, p := range c.populators {
		p.Close()
	}
}
