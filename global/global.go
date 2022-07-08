package global

import (
	"embed"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

var (
	// Conf 系统配置
	Conf Configuration
	// ConfBox 打包配置文件到golang编译后的二进制程序中
	ConfBox *CustomConfBox
	// Log zap日志
	Log *zap.SugaredLogger
	// Consul consul做服务发现
	Consul *consulapi.Client
)

// CustomConfBox 自定义配置盒子
type CustomConfBox struct {
	// 命令行指定的配置文件路径
	ConfigFile string
	// 丢弃packr盒子，使用go1.16新特性embed来嵌入配置文件到二进制执行文件中
	EmbedFS *embed.FS
	// viper实例
	ViperIns *viper.Viper
}

// Find 查找指定配置
func (c *CustomConfBox) Find(filename string) []byte {
	bs, _ := os.ReadFile(filename)
	if len(bs) == 0 {
		bs, _ = c.EmbedFS.ReadFile(filename)
	}
	return bs
}
