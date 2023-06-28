package test

import (
	"fmt"
	"testing"

	"github.com/ribincao/ribin-dev-box/ribin-common/config"
)

func TestConfig(t *testing.T) {
	config.InitConfig("./conf.yaml")
	fmt.Println("ConfigTest Env:", config.GlobalConfig.ServiceConfig.Env)
}
