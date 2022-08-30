package main

import (
	"MQP/config"
	"MQP/internal/populator"
)

func main() {
	endChan := make(chan bool)
	config.Init()
	controller := populator.Controller{}
	controller.Init()
	controller.Start()
	<-endChan
}
