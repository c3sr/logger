package logger

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

// AWS holds common AWS credentials and keys.
type loggerConfig struct {
	Level string   `json:"level" config:"logger.level"`
	Hooks []string `json:"hooks" config:"logger.hooks" default:'["syslog"]'`
}

var (
	Config = &loggerConfig{}
)

func (loggerConfig) ConfigName() string {
	return "AWS"
}

func (loggerConfig) SetDefaults() {
}

func (a *loggerConfig) Read() {
	vipertags.Fill(a)
}

func (c loggerConfig) String() string {
	return pp.Sprintln(c)
}

func (c loggerConfig) Debug() {
	Debug("AWS Config = ", c)
}

func init() {
	config.Register(Config)
}
