package populator

import (
	"MQP/config"
	"MQP/providers"
	"fmt"
	"time"
)

type Populator struct {
	provider     providers.Provider
	queuesConfig []config.Queue
}

func (p *Populator) Close() {
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

func (p Populator) getBody(queueConfig config.Queue) string {
	if queueConfig.Message.IncludeTimestamp {
		return fmt.Sprintf("%s %s", queueConfig.Message.Body, time.Now().String())
	} else {
		return queueConfig.Message.Body
	}
}

func (p Populator) PushMessages(queueConfig config.Queue) {
	for {
		delay := p.getDelay(queueConfig)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		body := p.getBody(queueConfig)
		p.provider.PushMessage(queueConfig.Name, body)
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
