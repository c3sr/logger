package hooks

import (
	"github.com/knq/sdhook"
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
)

func init() {
	config.AfterInit(func() {
		h, err := sdhook.New(
			sdhook.GoogleLoggingAgent(),
			sdhook.ErrorReportingService(config.App.Name),
		)
		if err != nil {
			return
		}
		logger.RegisterHook("stackdriver", h)
	})
}
