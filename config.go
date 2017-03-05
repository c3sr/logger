package logger

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type loggerConfig struct {
	Level      string   `json:"level" config:"logger.level"`
	Stacktrace bool     `json:"stack_trace" config:"logger.stack_trace" default:"true"`
	Hooks      []string `json:"hooks" config:"logger.hooks" default:'["syslog"]'`
}

var (
	Config = &loggerConfig{}
)

func (loggerConfig) ConfigName() string {
	return "Logger"
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
	Debug("Logger Config = ", c)
}

func init() {
	config.Register(Config)
}
