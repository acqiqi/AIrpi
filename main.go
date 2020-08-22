package main

import (
	"AIrpi/platforms/mqtt"
	"log"
)

func main() {
	log.Println("Ojbk")

	mt := mqtt.NewAdaptor("tcp://120.77.223.157:1883", "cccccclient")
	if err := mt.Connect(); err != nil {
		log.Println(err.Error())
	}
	str := "hahaha"
	mt.Publish("test", []byte(str))
	mt.On("/#", func(msg mqtt.Message) {
		topic := msg.Topic()
		log.Println(topic)
		log.Println(string(msg.Payload()))
	})

	select {}
}
