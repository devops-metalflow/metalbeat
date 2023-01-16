package initialize

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/spf13/viper"
	"metalbeat/global"
	"metalbeat/utils"
	"strings"
)

const (
	configType             = "yml"
	developmentConfig      = "conf/config.yml"
	defaultTimeout         = 5
	defaultDeregisterAfter = 30
)

// go 1.16新特性embed的嵌入文件是相对于当前源码文件所在目录的
//
//go:embed conf/config.yml
var f embed.FS

// Config 初始化配置文件
func Config(filename string) {
	// 初始化配置盒子
	var box global.CustomConfBox
	if filename != "" {
		if strings.HasPrefix(filename, "/") {
			// 指定的目录为绝对路径
			box.ConfigFile = filename
		} else {
			// 指定的目录为相对路径
			box.ConfigFile = utils.GetWorkDir() + "/" + filename
		}
	}
	// 获取viper实例
	box.ViperIns = viper.New()
	box.EmbedFS = &f
	global.ConfBox = &box
	v := box.ViperIns

	// 读取默认的配置文件作为默认配置项
	readConfig(v, developmentConfig)
	// 将默认的配置文件中的配置全部以默认配置写入
	settings := v.AllSettings()
	for index, setting := range settings {
		v.SetDefault(index, setting)
	}
	// 读取命令行配置文件内容
	if box.ConfigFile != "" {
		// 读取不同配置文件中的差异部分
		readConfig(v, box.ConfigFile)
	}
	// 转换为结构体
	if err := v.Unmarshal(&global.Conf); err != nil {
		panic(fmt.Sprintf("初始化配置文件失败: %v, 命令行配置文件路径: %s", err, global.ConfBox.ConfigFile))
	}

	// 健康检查的链接后缀检查: 命令服务的链接后缀检查
	if !strings.HasPrefix(global.Conf.Service.CheckSuffix, "/") {
		global.Conf.Service.CheckSuffix = fmt.Sprintf("/%s", global.Conf.Service.CheckSuffix)
	}
	if !strings.HasPrefix(global.Conf.Shell.ShellSuffix, "/") {
		global.Conf.Shell.ShellSuffix = fmt.Sprintf("/%s", global.Conf.Shell.ShellSuffix)
	}

	if global.Conf.Shell.Timeout < 1 {
		global.Conf.Shell.Timeout = defaultTimeout
	}
	if global.Conf.Service.CheckTimeout < 1 {
		global.Conf.Service.CheckTimeout = defaultTimeout
	}
	if global.Conf.Service.CheckInterval < 1 {
		global.Conf.Service.CheckInterval = defaultTimeout
	}
	if global.Conf.Service.DeregisterAfter < 1 {
		global.Conf.Service.DeregisterAfter = defaultDeregisterAfter
	}

	fmt.Println("初始化配置文件完成, 配置文件: ", global.ConfBox.ConfigFile)
}

func readConfig(v *viper.Viper, configFile string) {
	v.SetConfigType(configType)
	config := global.ConfBox.Find(configFile)
	if len(config) == 0 {
		panic(fmt.Sprintf("初始化配置文件失败: %v", configFile))
	}
	// 加载配置
	if err := v.ReadConfig(bytes.NewReader(config)); err != nil {
		panic(fmt.Sprintf("加载配置文件失败: %v", err))
	}
}
