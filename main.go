package main

import (
	"fmt"
	"rule_engine_by_go/utils"
	"rule_engine_by_go/utils/log"
)

func main() {

	config, _ := utils.GetConfig()
	log.Info.Println("test")
	fmt.Println(config)
}
