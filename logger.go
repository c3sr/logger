package logger

import (
	log "github.com/Sirupsen/logrus"
	"github.com/fatih/color"

	"github.com/rai-project/config"
)

type Logger struct {
	*log.Logger
	mu MutexWrap
}

var (
	debug    = false
	addhooks = true
	std      = New()
)

func setupHooks(log *Logger) {
	if addhooks {
		setupSyslogHook(log)
	}
}

func init() {
	config.OnInit(func() {
		formatter := &log.TextFormatter{
			DisableColors:    color.NoColor,
			ForceColors:      !color.NoColor,
			DisableTimestamp: true,
		}
		log.SetFormatter(formatter)
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
		setupHooks(&Logger{Logger: log.StandardLogger()})
		setupHooks(std)
	})
}
