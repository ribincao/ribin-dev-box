package models

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
