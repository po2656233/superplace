package superGORM

import (
	"strings"

	clog "github.com/po2656233/superplace/logger"
)

type gormLogger struct {
	log *clog.SuperLogger
}

func (l gormLogger) Printf(s string, i ...interface{}) {
	l.log.Debugf(strings.ReplaceAll(s, "\n", ""), i...)
}
