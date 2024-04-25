package superCron

import (
	"testing"
	"time"

	ctime "github.com/po2656233/superplace/extend/time"
	clog "github.com/po2656233/superplace/logger"
)

func TestAddEveryDayFunc(t *testing.T) {
	AddEveryDayFunc(func() {
		clog.Info(ctime.Now().ToDateTimeFormat())
	}, 17, 32, 5)

	AddEveryHourFunc(func() {
		clog.Info(ctime.Now().ToDateTimeFormat())
		panic("print panic~~~")
	}, 5, 5)

	AddDurationFunc(func() {
		clog.Info(ctime.Now().ToDateTimeFormat())
	}, 1*time.Minute)

	Run()
}
