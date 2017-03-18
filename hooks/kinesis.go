package hooks

import (
	"github.com/evalphobia/logrus_kinesis"
	"github.com/rai-project/aws"
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
)

func init() {
	config.OnInit(func() {
		logger.Config.Wait()

		if !logger.UsingHook("kinesis") {
			return
		}

		aws.Config.Wait()

		c, err := aws.NewConfig()
		if err != nil {
			return
		}
		h, err := logrus_kinesis.NewWithAWSConfig(config.App.Name, c)
		if err != nil {
			return
		}
		logger.RegisterHook("kinesis", h)
	})
}
