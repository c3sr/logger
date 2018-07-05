package hooks

import (
	"os"

	"github.com/rai-project/config"
	"github.com/rai-project/logger"
	"github.com/rai-project/utils"
	"github.com/spf13/viper"
	"gopkg.in/gemnasium/logrus-airbrake-hook.v3"
)

func init() {
	config.OnInit(func() {
		config.App.Wait()
		logger.Config.Wait()

		if !logger.UsingHook("airbrake") {
			return
		}

		projectId := viper.GetInt64("airbrake.id")
		if projectId == 0 {
			return
		}
		apiKey := decrypt(viper.GetString("airbrake.api_key"))
		if apiKey == "" {
			return
		}

		env := "ID=" + config.App.Name + ";" +
			"Version=" + config.App.Version.Version + ";" +
			"BuildDate=" + config.App.Version.BuildDate

		if ip, err := utils.GetExternalIp(); err == nil {
			env = env + ";IP=" + ip
		}

		if hostname, err := os.Hostname(); err == nil {
			env = env + ";HostName=" + hostname
		}

		hook := airbrake.NewHook(projectId, apiKey, env)

		logger.RegisterHook("airbrake", hook)
	})
}
