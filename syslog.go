// +build !windows

package logger

import (
	"log/syslog"

	logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
)

func setupSyslogHook(log *Logger) {
	hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_DEBUG, "")
	if err != nil {
		return
	}
	log.Hooks.Add(hook)
}
