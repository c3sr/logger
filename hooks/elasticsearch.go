// +build disabled

package hooks

import (
	"fmt"

	"github.com/c3sr/config"
	"github.com/c3sr/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	config.OnInit(func() {
		logger.Config.Wait()

		if !logger.UsingHook("elasticsearch") {
			return
		}

		logger.Config.Wait()

		host := viper.GetString("elasticsearch.host")
		port := viper.GetInt("elasticsearch.port")
		index := viper.GetString("elasticsearch.index")
		level, err := logrus.ParseLevel(logger.Config.Level)
		if err != nil {
			fmt.Println("cannot parse level from logger config ", err)
			return
		}

		client, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%s:%d", host, port)))
		if err != nil {
			fmt.Println("failed to load elasticsearch client for logger hook ", err)
			return
		}
		h, err := elogrus.NewElasticHook(client, host, level, index)
		if err != nil {
			fmt.Println("failed to load elasticsearch logger hook ", err)
			return
		}

		logger.RegisterHook("elasticsearch", h)
	})
}
