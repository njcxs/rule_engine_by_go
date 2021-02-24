package app

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
