// +build linux

package hooks

import (
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
	"github.com/wercker/journalhook"
)

func init() {
	config.OnInit(func() {
		h := &journalhook.JournalHook{}
		logger.RegisterHook("journald", h)
	})
}
