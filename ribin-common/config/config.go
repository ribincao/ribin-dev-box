package config

import (
	"io/ioutil"

	"github.com/ribincao/ribin-dev-box/ribin-common/constant"
	"gopkg.in/yaml.v2"
)

var GlobalConfig *Config

func InitConfig(path string) {
	if path == "" {
		path = constant.DEFAULT_CONFIG_PATH
	}
	GlobalConfig = parseConfig(path)
}

func parseConfig(configPath string) *Config {
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	config := &Config{}
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		panic(err)
	}

	return config
}

type Config struct {
	ServiceConfig *ServiceConfig  `yaml:"serviceConfig"`
	LogConfig     *LogConfig      `yaml:"logConfig"`
	AgonesConfig  []*AgonesConfig `yaml:"agonesConfig"`
}

type LogConfig struct {
	LogPath    string `yaml:"logPath"`
	LogLevel   string `yaml:"logLevel"`
	LogMaxAge  int    `yaml:"logMaxAge"`
	LogMaxSize int    `yaml:"logMaxSize"`
	LogMode    string `yaml:"logMode"`
}

type ServiceConfig struct {
	Env           string `yaml:"env"`
	Port          string `yaml:"port"`
	RedisAddr     string `yaml:"redisAddr"`
	RedisUserName string `yaml:"redisUserName"`
	RedisPasswd   string `yaml:"redisPasswd"`
}

type AgonesConfig struct {
	Region       string `yaml:"region"`
	Prefix       string `yaml:"prefix"`
	AgonesAddr   string `yaml:"agonesAddr"`
	NameSpace    string `yaml:"nameSpace"`
	MultiCluster bool   `yaml:"multiCluster"`
}
