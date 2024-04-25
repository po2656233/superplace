package superCron

import (
	exReflect "github.com/po2656233/superplace/extend/reflect"
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	"github.com/robfig/cron/v3"
)

type Component struct {
	cfacade.Component
}

// Name unique components name
func (*Component) Name() string {
	return exReflect.GetPackName(Component{})
}

func (p *Component) Init() {
	Start()
	clog.Info("cron component init.")
}

func (p *Component) OnStop() {
	Stop()
	clog.Infof("cron component is stopped.")
}

func New(opts ...cron.Option) cfacade.IComponent {
	if len(opts) > 0 {
		Init(opts...)
	}
	return &Component{}
}
