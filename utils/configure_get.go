package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var configFile []byte

func GetConfig() (result map[string]interface{}, err error) {
	err = yaml.Unmarshal(configFile, &result)
	return result, err
}

func init() {
	var err error
	configFile, err = ioutil.ReadFile(GetCurrentPath() + "/etc/config.yml")
	if err != nil {
		log.Fatalf("Get yml file err %v ", err)
	}

}
