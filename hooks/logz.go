package hooks

import (
	"fmt"

	"github.com/bshuster-repo/logruzio"
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
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

		token := viper.GetString("logz.token")

		ctx := logrus.Fields{
			"ID":        config.App.Name,
			"Version":   config.App.Version.Version,
			"BuildDate": config.App.Version.BuildDate,
		}
		hook, err := logruzio.New(token, config.App.Name, ctx)
		if err != nil {
			fmt.Println("cannot register logz hook ", err)
			return
		}

		logger.RegisterHook("logz", hook)
	})
}
