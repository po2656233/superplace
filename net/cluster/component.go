package cluster

import (
	exReflect "github.com/po2656233/superplace/extend/reflect"
	face "github.com/po2656233/superplace/facade"
	superNatsCluster "github.com/po2656233/superplace/net/cluster/nats_cluster"
)

type Component struct {
	face.Component
	face.ICluster
}

func New() *Component {
	return &Component{}
}

func (c *Component) Name() string {
	return exReflect.GetPackName(Component{})
}

func (c *Component) Init() {
	c.ICluster = c.loadCluster()
	c.ICluster.Init()
}

func (c *Component) OnStop() {
	c.ICluster.Stop()
}

func (c *Component) loadCluster() face.ICluster {
	return superNatsCluster.New(c.App())
}
