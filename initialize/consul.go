package initialize

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"metalbeat/global"
)

// Consul 初始化consul
func Consul() {
	config := consulapi.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", global.Conf.Consul.Address, global.Conf.Consul.Port)
	client, err := consulapi.NewClient(config)
	if err != nil {
		panic(fmt.Sprintf("初始化consul异常:%v", err))
	}

	// check agent valid.
	_, err = client.Agent().Checks()
	if err != nil {
		panic(fmt.Sprintf("连接consul异常: %v", err))
	}

	global.Consul = client
	global.Log.Info("初始化consul完成")
}
