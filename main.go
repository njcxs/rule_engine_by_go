//package main
//
//import (
//	"fmt"
//	"rule_engine_by_go/utils"
//	"rule_engine_by_go/utils/log"
//)
//
//func main() {
//
//	config, _ := utils.GetConfig()
//	log.Info.Println("test")
//	fmt.Println(config["input"])
//}

package main

import "rule_engine_by_go/utils/json"

const json_ = `{"name":{"first":"Janet","last":"Prichard"},".age":47}`

func main() {
	value := json.Get(json_, "name.last")
	println(value.String())
}
