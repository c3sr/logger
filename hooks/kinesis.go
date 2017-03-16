package hooks

import (
	"github.com/evalphobia/logrus_kinesis"
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
)

func init() {
	config.AfterInit(func() {
		h, err := logrus_kinesis.New(config.App.Name, Config{
			AccessKey: "ABC", // AWS accessKeyId
			SecretKey: "XYZ", // AWS secretAccessKey
			Region:    "ap-northeast-1",
		})
		if err != nil {
			return
		}
		logger.RegisterHook("kinesis", h)
	})
}
