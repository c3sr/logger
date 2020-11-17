// +build linux

package hooks

import (
	"github.com/c3sr/config"
	"github.com/c3sr/logger"
	"github.com/wercker/journalhook"
)

func init() {
	config.OnInit(func() {
		logger.Config.Wait()

		if !logger.UsingHook("journald") {
			return
		}

		h := &journalhook.JournalHook{}
		logger.RegisterHook("journald", h)
	})
}
