package hooks

import (
	"fmt"
	"os"
	"time"

	"github.com/rai-project/config"
	"github.com/rai-project/logger"
	"github.com/rai-project/logger/hooks/logruzio"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	config.OnInit(func() {
		config.App.Wait()
		logger.Config.Wait()

		if !logger.UsingHook("logz") {
			return
		}

		token := decrypt(viper.GetString("logz.token"))
		if token == "" {
			return
		}

		ctx := logrus.Fields{
			"ID":        config.App.Name,
			"Version":   config.App.Version.Version,
			"BuildDate": config.App.Version.BuildDate,
		}
		if hostname, err := os.Hostname(); err == nil {
			ctx["HostName"] = hostname
		}
		hook, err := logruzio.New(token, config.App.Name, 5*time.Minute, ctx)
		if err != nil {
			fmt.Println("cannot register logz hook ", err)
			return
		}

		logger.RegisterHook("logz", hook)
	})
}
