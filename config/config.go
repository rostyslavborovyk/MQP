package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type RandomConfig struct {
	ErlangOrder int `json:"erlangOrder"`
}

type Message struct {
	Body             string       `json:"body"`
	IncludeTimestamp bool         `json:"includeTimestamp"`
	IncludeRandom    bool         `json:"includeRandom"`
	Frequency        float64      `json:"frequency"`
	RandomConfig     RandomConfig `json:"randomConfig"`
}

type Queue struct {
	Name    string  `json:"name"`
	Message Message `json:"message"`
}

type Service struct {
	Type   string  `json:"type"`
	Queues []Queue `json:"queues"`
}

type MQPConfig struct {
	Services []Service `json:"services"`
}

var config *MQPConfig

func Init() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Panicln("Unable to read config file")
	}

	config = &MQPConfig{}

	if err := json.Unmarshal([]byte(file), &config); err != nil {
		log.Panicf("Unable to parse config file %s", err)
	}
}

func GetConfig() *MQPConfig {
	if config == nil {
		log.Panicln("Config was not initialized")
	}
	return config
}
