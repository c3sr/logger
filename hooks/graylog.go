package hooks

import (
	"os"

	"github.com/rai-project/config"
	"github.com/rai-project/logger"
	"github.com/rai-project/utils"
	"github.com/spf13/viper"
	"gopkg.in/gemnasium/logrus-graylog-hook.v2"
)

func init() {
	config.OnInit(func() {
		config.App.Wait()
		logger.Config.Wait()

		if !logger.UsingHook("graylog") {
			return
		}

		address := decrypt(viper.GetString("graylog.address"))
		if address == "" {
			return
		}

		port := decrypt(viper.GetString("graylog.port"))
		if port == "" {
			port = "12201"
		}

		ctx := map[string]interface{}{
			"ID":        config.App.Name,
			"Version":   config.App.Version.Version,
			"BuildDate": config.App.Version.BuildDate,
		}

		if ip, err := utils.GetExternalIp(); err == nil {
			ctx["IP"] = ip
		}

		if hostname, err := os.Hostname(); err == nil {
			ctx["HostName"] = hostname
		}

		hook := graylog.NewGraylogHook(address+":"+port, ctx)

		logger.RegisterHook("graylog", hook)
	})
}
