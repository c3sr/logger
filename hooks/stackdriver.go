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
		found := false
		for _, h := range logger.Config.Hooks {
			if h == "stackdriver" {
				found = true
				break
			}
		}
		if !found {
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
