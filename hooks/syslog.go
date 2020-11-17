// +build !windows

package hooks

import (
	"log/syslog"

	"github.com/c3sr/config"
	"github.com/c3sr/logger"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
)

func init() {
	config.OnInit(func() {
		logger.Config.Wait()

		if !logger.UsingHook("syslog") {
			return
		}

		h, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_DEBUG, "")
		if err != nil {
			return
		}
		logger.RegisterHook("syslog", h)
	})
}
