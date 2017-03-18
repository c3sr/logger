// +build linux

package hooks

import (
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
	"github.com/wercker/journalhook"
)

func init() {
	config.OnInit(func() {
		logger.Config.Wait()
		found := false
		for _, h := range logger.Config.Hooks {
			if h == "journald" {
				found = true
				break
			}
		}
		if !found {
			return
		}

		h := &journalhook.JournalHook{}
		logger.RegisterHook("journald", h)
	})
}
