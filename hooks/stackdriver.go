package hooks

import (
	"github.com/knq/sdhook"
	"github.com/rai-project/config"
	"github.com/rai-project/googlecloud"
	"github.com/rai-project/logger"
)

func init() {
	config.OnInit(func() {
		logger.Config.Wait()

		if !logger.UsingHook("stackdriver") {
			return
		}

		googlecloud.Config.Wait()

		opts := googlecloud.NewOptions()

		h, err := sdhook.New(
			sdhook.GoogleLoggingAgent(),
			sdhook.GoogleServiceAccountCredentialsJSON(opts.Bytes()),
			sdhook.ErrorReportingService(config.App.Name),
		)
		if err != nil {
			return
		}
		logger.RegisterHook("stackdriver", h)
	})
}
