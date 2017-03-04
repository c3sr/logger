package logger

import (
	log "github.com/Sirupsen/logrus"
	"github.com/fatih/color"
	"github.com/spf13/viper"

	"github.com/rai-project/config"
)

type Logger struct {
	*log.Logger
	mu MutexWrap
}

var (
	debug = false
	std   = New()
)

func setupHooks(log *Logger) {
	for _, hook := range Config.Hooks {
		switch hook {
		case "syslog":
			setupSyslogHook(log)
		}
	}
}

func init() {
	config.OnInit(func() {
		formatter := &log.TextFormatter{
			DisableColors:    !viper.GetBool("color"),
			ForceColors:      viper.GetBool("color"),
			DisableTimestamp: true,
		}
		log.SetFormatter(formatter)
		log.SetOutput(color.Output)
		std.Formatter = formatter

		if config.IsVerbose {
			log.SetLevel(log.DebugLevel)
			std.Level = log.DebugLevel
		} else if config.IsDebug {
			log.SetLevel(log.DebugLevel)
			std.Level = log.DebugLevel
		} else {
			log.SetLevel(log.InfoLevel)
			std.Level = log.InfoLevel
		}

		if lvl, err := log.ParseLevel(Config.Level); err == nil {
			log.SetLevel(lvl)
			std.Level = lvl
		}

		setupHooks(&Logger{Logger: log.StandardLogger()})
		setupHooks(std)
	})
}
