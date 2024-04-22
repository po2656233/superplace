package extendCluster

import (
	face "superplace/facade"
	cherryNatsCluster "superplace/net/cluster/nats_cluster"
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
