package logger

import (
	"github.com/Sirupsen/logrus"
	"github.com/fatih/color"
	"github.com/spf13/viper"

	"github.com/rai-project/config"
)

type Logger struct {
	*logrus.Logger
	mu MutexWrap
}

var (
	debug = false
	std   = New()
)

func UsingHook(s string) bool {
	for _, h := range Config.Hooks {
		if h == s {
			return true
		}
	}
	return false
}

func setupHooks(log *Logger) {
	if UsingHook("stacktrace") && log.Level >= logrus.DebugLevel {
		log.Hooks.Add(StandardStackHook())
	}
	for _, h := range hooks.data {
		log.Hooks.Add(h)
	}
}

func init() {
	config.OnInit(func() {
		formatter := &logrus.TextFormatter{
			DisableColors:    !viper.GetBool("color"),
			ForceColors:      viper.GetBool("color"),
			DisableSorting:   true,
			DisableTimestamp: true,
		}
		logrus.SetFormatter(formatter)
		logrus.SetOutput(color.Output)
		std.Formatter = formatter

		if config.IsVerbose {
			logrus.SetLevel(logrus.DebugLevel)
			std.Level = logrus.DebugLevel
		} else if config.IsDebug {
			logrus.SetLevel(logrus.DebugLevel)
			std.Level = logrus.DebugLevel
		} else {
			logrus.SetLevel(logrus.InfoLevel)
			std.Level = logrus.InfoLevel
		}

		if lvl, err := logrus.ParseLevel(Config.Level); err == nil {
			logrus.SetLevel(lvl)
			std.Level = lvl
		}

		setupHooks(&Logger{Logger: logrus.StandardLogger()})
		setupHooks(std)
	})
}
