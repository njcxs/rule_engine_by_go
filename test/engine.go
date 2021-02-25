package main

import (
	"rule_engine_by_go/app"
	"rule_engine_by_go/utils/kafka"

	"fmt"
	"time"
)

func main() {

	e := app.NewEngine("conn")
	e.ReadRules()

	kafkaConsumer := kafka.InitKakfaConsumer([]string{"172.21.129.2:9092"}, "test", []string{"nids-conn"})
	kafkaConsumer.Open()

	go func() {

		for {

			message := <-kafkaConsumer.Message
			fmt.Println(string(message.Value))
		}
	}()

	go func() {

		for {

			message := <-kafkaConsumer.Message
			fmt.Println(string(message.Value))
		}
	}()

	go func() {

		for {

			message := <-kafkaConsumer.Message
			fmt.Println(string(message.Value))
		}
	}()

	go func() {

		for {

			message := <-kafkaConsumer.Message
			fmt.Println(string(message.Value))
		}
	}()

	go e.ResCheck()

	time.Sleep(10 * time.Second)
	//value, _ :=
	//fmt.Println(value["name"])
}
