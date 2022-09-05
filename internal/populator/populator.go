package populator

import (
	"fmt"
	"github.com/rostyslavborovyk/MQP/config"
	"github.com/rostyslavborovyk/MQP/pkg/providers"
	"math/rand"
	"time"
)

type Populator struct {
	provider     providers.Provider
	queuesConfig []config.Queue
}

func (p Populator) Close() {
	p.provider.Close()
}

func (p Populator) getDelay(queueConfig config.Queue) float64 {
	if queueConfig.Message.IncludeRandom {
		return GeneratePoissonInterval(
			queueConfig.Message.Frequency,
			queueConfig.Message.RandomConfig.ErlangOrder,
		)
	} else {
		return 1000 / queueConfig.Message.Frequency
	}
}

func getRandomVariation(variations []interface{}) interface{} {
	return variations[int(rand.Float64()*float64(len(variations)))]
}

func (p Populator) getBody(queueConfig config.Queue) interface{} {
	if queueConfig.Message.IncludeTimestamp && queueConfig.Message.BodyVariations.Type == "text/plain" {
		variation := getRandomVariation(queueConfig.Message.BodyVariations.Variations)
		return fmt.Sprintf("%s %s", variation, time.Now().String())
	} else {
		return getRandomVariation(queueConfig.Message.BodyVariations.Variations)
	}
}

func (p Populator) pushMessage(queueConfig config.Queue) {
	delay := p.getDelay(queueConfig)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	body := p.getBody(queueConfig)
	p.provider.PushMessage(
		queueConfig.Name,
		queueConfig.Message.BodyVariations.Type,
		body,
	)
}

func (p Populator) PushMessages(queueConfig config.Queue) {
	for {
		p.pushMessage(queueConfig)
	}
}

func (p Populator) RunQueue(queueConfig config.Queue) {
	p.provider.CreateQueue(queueConfig.Name)
	go func() {
		p.PushMessages(queueConfig)
	}()
}

func (p Populator) Run() {
	for _, qc := range p.queuesConfig {
		p.RunQueue(qc)
	}
}
