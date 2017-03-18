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
		found := false
		for _, h := range logger.Config.Hooks {
			if h == "kinesis" {
				found = true
				break
			}
		}
		if !found {
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
