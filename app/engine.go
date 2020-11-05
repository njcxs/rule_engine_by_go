package app

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/flier/gohs/hyperscan" // 引入
	"github.com/gomodule/redigo/redis"
	"github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
	"rule_engine_by_go/utils/kafka"
	"rule_engine_by_go/utils/log"
	"rule_engine_by_go/utils/workerpool"
	"strings"
)

type engine struct {
	RuleType string
}

func NewEngine(ruleType string) *engine {
	return &engine{RuleType: ruleType}
}

func (e *engine) ReadRules() {

}

func (e *engine) ResCheck() {
}

func (e *engine) Input() {

}

func (e *engine) OutPut() {

}
