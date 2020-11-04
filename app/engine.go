package app

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gomodule/redigo/redis"
	"github.com/json-iterator/go" // 引入
	"rule_engine_by_go/utils/kafka"
	"rule_engine_by_go/utils/log"
	"rule_engine_by_go/utils/workerpool"
	"strings"
)

type engine struct {
	rule_type string
}

func NewEngine(rule_type string) *engine {

}

func (e *engine) ReadRules() {

}
