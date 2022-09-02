package main

import (
	"github.com/rostyslavborovyk/MQP/config"
	"github.com/rostyslavborovyk/MQP/internal/populator"
)

func main() {
	endChan := make(chan bool)
	config.Init()
	controller := populator.Controller{}
	controller.Init()
	controller.Start()
	<-endChan
}
