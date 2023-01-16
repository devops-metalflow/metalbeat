package global

import "go.uber.org/zap/zapcore"

// Configuration 系统配置，配置字段可参考yml注释
// viper内置了mapstructure， yml文件用“-”区分单词，转为驼峰方便
type Configuration struct {
	System  SystemConfiguration  `mapstructure:"system" json:"system"`
	Shell   ShellConfiguration   `mapstructure:"shell" json:"shell"`
	Service ServiceConfiguration `mapstructure:"service" json:"service"`
	Consul  ConsulConfiguration  `mapstructure:"consul" json:"consul"`
	Logs    LogsConfiguration    `mapstructure:"logs" json:"logs"`
}

type SystemConfiguration struct {
	Kind     string `mapstructure:"kind" json:"kind"`
	Metadata string `mapstructure:"metadata" json:"metadata"`
}

type ShellConfiguration struct {
	ShellSuffix  string                   `mapstructure:"shell-suffix" json:"shellSuffix"`
	Timeout      int                      `mapstructure:"timeout" json:"timeout"`
	InitShell    bool                     `mapstructure:"init-shell" json:"initShell"`
	InitCommands []MetalTaskConfiguration `mapstructure:"init-commands" json:"initCommands"`
}

type MetalTaskConfiguration struct {
	Name        string   `mapstructure:"name" json:"name"`
	Annotations string   `mapstructure:"annotations" json:"annotations"`
	Commands    []string `mapstructure:"commands" json:"commands"`
}

type ServiceConfiguration struct {
	Name            string `mapstructure:"name" json:"name"`
	Tags            string `mapstructure:"tags" json:"tags"`
	Port            int    `mapstructure:"port" json:"port"`
	CheckTimeout    int    `mapstructure:"check-timeout" json:"checkTimeout"`
	CheckSuffix     string `mapstructure:"check-suffix" json:"checkSuffix"`
	CheckInterval   int    `mapstructure:"check-interval" json:"checkInterval"`
	DeregisterAfter int    `mapstructure:"deregister-after" json:"deregisterAfter"`
}

type ConsulConfiguration struct {
	Address string `mapstructure:"address" json:"address"`
	Port    int    `mapstructure:"port" json:"port"`
}

type LogsConfiguration struct {
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}
