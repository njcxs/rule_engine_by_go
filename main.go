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

import (
	"fmt"
	"rule_engine_by_go/app"
)

//
//const json_ = `{"name":{"first":"Janet","last":"Prichard"},".age":47}`

func main() {
	value, _ := app.GetConfig()
	fmt.Println(value["name"])
}
