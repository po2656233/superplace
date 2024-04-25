package superplace

import (
	"github.com/po2656233/superplace/extend/proc"
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	ccluster "github.com/po2656233/superplace/net/cluster"
	cdiscovery "github.com/po2656233/superplace/net/discovery"
	"github.com/po2656233/superplace/tools"
)

type (
	AppBuilder struct {
		*Application
		components []cfacade.IComponent
	}
)

func Configure(profileFilePath, nodeId string, isFrontend bool, mode NodeMode) *AppBuilder {
	appBuilder := &AppBuilder{
		Application: NewApp(profileFilePath, nodeId, isFrontend, mode),
		components:  make([]cfacade.IComponent, 0),
	}

	return appBuilder
}

func ConfigureNode(node cfacade.INode, isFrontend bool, mode NodeMode) *AppBuilder {
	appBuilder := &AppBuilder{
		Application: NewAppNode(node, isFrontend, mode),
		components:  make([]cfacade.IComponent, 0),
	}

	return appBuilder
}

func (p *AppBuilder) Startup() {
	app := p.Application

	if app.NodeMode() == Cluster {
		cluster := ccluster.New()
		app.SetCluster(cluster)
		app.Register(cluster)

		discovery := cdiscovery.New()
		app.SetDiscovery(discovery)
		app.Register(discovery)

		// 启动nats
		if ok, _ := proc.CheckProcRunning("nats-server"); !ok {
			err := proc.StartCMD(tools.GetNatsSHFile())
			clog.Infof("nats-server START err:%v", err)
		}
	}

	// Register custom components
	app.Register(p.components...)

	// startup
	app.Startup()
}

func (p *AppBuilder) Register(component ...cfacade.IComponent) {
	p.components = append(p.components, component...)
}

func (p *AppBuilder) AddActors(actors ...cfacade.IActorHandler) {
	p.actorSystem.Add(actors...)
}

func (p *AppBuilder) NetParser() cfacade.INetParser {
	return p.netParser
}

func (p *AppBuilder) SetNetParser(parser cfacade.INetParser) {
	p.netParser = parser
}
