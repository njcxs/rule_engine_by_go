package app

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"rule_engine_by_go/utils"
	log2 "rule_engine_by_go/utils/log"
)

type engine struct {
	RuleType string
	InPutC   chan string
	OutPutC  chan string
	RuleList []map[interface{}]interface{}
}

func NewEngine(ruleType string) *engine {
	return &engine{RuleType: ruleType}
}

func (e *engine) ReadRules() {
	var rulesPath []string
	rulesListPath, err := utils.GetAllFile(utils.GetCurrentPath()+"/etc/rules/"+e.RuleType, rulesPath)
	if err != nil {
		log2.Error.Fatalf("Get rule file dir err %s  %v ", rulesListPath, err)
	}

	for _, ruleFile := range rulesListPath {
		rule, err := ioutil.ReadFile(ruleFile)

		if err != nil {
			log2.Error.Fatalf("Get rule file err  %s %v ", rule, err)
		}
		m := make(map[interface{}]interface{})
		err = yaml.Unmarshal(rule, &m)
		if err != nil {
			log2.Error.Fatalf("Unmarshal rule file err  %s %v ", rule, err)
		}
		e.RuleList = append(e.RuleList, m)
	}

	log2.Info.Printf("Rule file have been load %s ", e.RuleType)

}

func (e *engine) ResCheck() {

	go func() {
		for {
			fmt.Println(<-e.InPutC)
		}

	}()

}

//func (e *engine) Input() {
//
//
//
//}
//
//func (e *engine) OutPut() {
//
//}
