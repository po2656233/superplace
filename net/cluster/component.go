package extendCluster

import (
	face "github.com/po2656233/superplace/facade"
	cherryNatsCluster "github.com/po2656233/superplace/net/cluster/nats_cluster"
)

const (
	Name = "cluster_component"
)

type Component struct {
	face.Component
	face.ICluster
}

func New() *Component {
	return &Component{}
}

func (c *Component) Name() string {
	return Name
}

func (c *Component) Init() {
	c.ICluster = c.loadCluster()
	c.ICluster.Init()
}

func (c *Component) OnStop() {
	c.ICluster.Stop()
}

func (c *Component) loadCluster() face.ICluster {
	return cherryNatsCluster.New(c.App())
}
