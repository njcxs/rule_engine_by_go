package main

import (
	"fmt"
	"rule_engine_by_go/utils/kafka"
)

func main() {

	kafkaConsumer := kafka.InitKakfaConsumer([]string{"10.10.128.235:9093"}, "test", []string{"hids"})
	kafkaConsumer.Open()

	for {

		message := <-kafkaConsumer.Message
		fmt.Println(string(message.Value))
	}

}
