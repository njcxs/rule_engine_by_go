package main

import (
	"rule_engine_by_go/utils/kafka"

	"fmt"
)

func main() {
	data := kafka.Data{Log: "test"}

	test := kafka.DataProducer{}
	test.Init([]string{"10.10.128.235:9093"}, "test", "test")
	fmt.Println(test.AddMessage(data))

}
