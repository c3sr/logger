package hooks

import (
	"strings"

	"github.com/rai-project/config"
	"github.com/rai-project/utils"
)

func decrypt(s string) string {
	config.App.Wait()
	if strings.HasPrefix(s, utils.CryptoHeader) && config.App.Secret != "" {
		if val, err := utils.DecryptStringBase64(config.App.Secret, s); err == nil {
			return val
		}
	}
	return s
}
