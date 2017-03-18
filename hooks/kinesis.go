package hooks

import (
	"github.com/apex/log/handlers/kinesis"
	"github.com/evalphobia/logrus_kinesis"
	"github.com/rai-project/aws"
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
)

func init() {
	config.OnInit(func() {
		aws.Config.Wait()

		sess, err := aws.NewSession()
		if err != nil {
			return
		}
		svc := kinesis.New(sess)
		h := &logrus_kinesis.KinesisHook{
			client:            svc,
			defaultStreamName: name,
			levels:            defaultLevels,
			ignoreFields:      make(map[string]struct{}),
			filters:           make(map[string]func(interface{}) interface{}),
		}
		logger.RegisterHook("kinesis", h)
	})
}
