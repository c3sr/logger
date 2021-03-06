package logger

import (
	"github.com/k0kubun/pp/v3"
	"github.com/c3sr/config"
	"github.com/c3sr/vipertags"
)

type loggerConfig struct {
	Level string        `json:"level" config:"logger.level" default:"info"`
	Hooks []string      `json:"hooks" config:"logger.hooks" default:'[]'`
	done  chan struct{} `json:"-" config:"-"`
}

var (
	Config = &loggerConfig{
		done: make(chan struct{}),
	}
)

func (loggerConfig) ConfigName() string {
	return "Logger"
}

func (a *loggerConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *loggerConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
}

func (c loggerConfig) Wait() {
	<-c.done
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
