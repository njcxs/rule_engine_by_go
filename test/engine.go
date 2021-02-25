package main

import (
	"rule_engine_by_go/app"
)

func main() {

	e := app.NewEngine("conn")
	e.ReadRules()
	//value, _ :=
	//fmt.Println(value["name"])
}
