package cherryDiscovery

import (
	cprofile "github.com/po2656233/superplace/config"
	face "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
)

const (
	Name = "discovery_component"
)

type Component struct {
	face.Component
	face.IDiscovery
}

func New() *Component {
	return &Component{}
}

func (*Component) Name() string {
	return Name
}

func (p *Component) Init() {
	config := cprofile.GetConfig("cluster").GetConfig("discovery")
	if config.LastError() != nil {
		clog.Error("`cluster` property not found in config file.")
		return
	}

	mode := config.GetString("mode")
	if mode == "" {
		clog.Error("`discovery->mode` property not found in config file.")
		return
	}

	discovery, found := discoveryMap[mode]
	if discovery == nil || !found {
		clog.Errorf("mode = %s property not found in discovery config.", mode)
		return
	}

	clog.Infof("Select discovery [mode = %s].", mode)
	p.IDiscovery = discovery
	p.IDiscovery.Load(p.App())
}

func (p *Component) OnStop() {
	p.IDiscovery.Stop()
}
