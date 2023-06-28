package config

import (
	"io/ioutil"

	"github.com/ribincao/ribin-dev-box/ribin-common/constant"
	"github.com/ribincao/ribin-dev-box/ribin-common/models"
	"gopkg.in/yaml.v2"
)

var GlobalConfig *models.Config

func InitConfig(path string) {
	if path == "" {
		path = constant.DEFAULT_CONFIG_PATH
	}
	GlobalConfig = parseConfig(path)
}

func parseConfig(configPath string) *models.Config {
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	config := &models.Config{}
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		panic(err)
	}

	return config
}
