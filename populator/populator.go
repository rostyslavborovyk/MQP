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

func (p Populator) PushMessagesWithDelay(queueConfig config.Queue) {
	delay := 1000 / queueConfig.Message.Frequency
	for {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		var body string
		if queueConfig.Message.IncludeTimestamp {
			body = fmt.Sprintf("%s %s", queueConfig.Message.Body, time.Now().String())
		} else {
			body = queueConfig.Message.Body
		}
		p.provider.PushMessage(queueConfig.Name, body)
	}
}

func (p Populator) RunQueue(queueConfig config.Queue) {
	p.provider.CreateQueue(queueConfig.Name)
	go func() {
		if !queueConfig.Message.IncludeRandom {
			p.PushMessagesWithDelay(queueConfig)
		}
	}()
}

func (p Populator) Run() {
	for _, qc := range p.queuesConfig {
		p.RunQueue(qc)
	}
}
