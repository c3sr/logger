// +build !windows

package hooks

import (
	"log/syslog"

	logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
)

func init() {
	config.OnInit(func() {
		logger.Config.Wait()
		found := false
		for _, h := range logger.Config.Hooks {
			if h == "syslog" {
				found = true
				break
			}
		}
		if !found {
			return
		}

		h, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_DEBUG, "")
		if err != nil {
			return
		}
		logger.RegisterHook("syslog", h)
	})
}
