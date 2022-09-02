package main

import (
	"github.com/rostyslavborovyk/config"
	"github.com/rostyslavborovyk/internal/populator"
)

func main() {
	endChan := make(chan bool)
	config.Init()
	controller := populator.Controller{}
	controller.Init()
	controller.Start()
	<-endChan
}
