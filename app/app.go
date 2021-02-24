package app

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"rule_engine_by_go/utils"
	log2 "rule_engine_by_go/utils/log"
)

var configFile []byte

func GetConfig() (result map[string]interface{}, err error) {
	err = yaml.Unmarshal(configFile, &result)
	return result, err
}

func init() {
	var err error
	configFile, err = ioutil.ReadFile(utils.GetCurrentPath() + "/etc/config.yml")
	if err != nil {
		log2.Error.Fatalf("Get yml file err %v ", err)
	}

}
