package logger

import (
	"github.com/Sirupsen/logrus"
	kinesis "github.com/evalphobia/logrus_kinesis"
	"github.com/rai-project/config"
	"github.com/spf13/viper"
)

func setupKenisisHook(log *Logger) {
	hook, err := kinesis.New(config.App.Name, kinesis.Config{
		AccessKey: viper.GetString("aws.access_key_id"),
		SecretKey: viper.GetString("aws.secret_access_key"),
		Region:    viper.GetString("aws.region"),
	})
	if err != nil {
		return
	}

	// set custom fire level
	hook.SetLevels([]logrus.Level{
		logrus.PanicLevel,
		logrus.ErrorLevel,
	})

	log.Hooks.Add(hook)
}
