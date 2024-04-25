package superGops

import (
	"github.com/google/gops/agent"
	exReflect "github.com/po2656233/superplace/extend/reflect"
	sgxFacade "github.com/po2656233/superplace/facade"
	sgxLogger "github.com/po2656233/superplace/logger"
)

// Component gops 监听进程数据
type Component struct {
	sgxFacade.Component
	options agent.Options
}

func New(options ...agent.Options) *Component {
	component := &Component{}
	if len(options) > 0 {
		component.options = options[0]
	}
	return component
}

func (c *Component) Name() string {
	return exReflect.GetPackName(Component{})
}

func (c *Component) Init() {
	if err := agent.Listen(c.options); err != nil {
		sgxLogger.Error(err)
	}
}

func (c *Component) OnAfterInit() {
}

func (c *Component) OnStop() {
	//agent.Close()
}
