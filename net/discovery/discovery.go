package cherryDiscovery

import (
	face "superplace/facade"
	clog "superplace/logger"
)

var (
	discoveryMap = make(map[string]face.IDiscovery)
)

func init() {
	Register(&DiscoveryDefault{})
	Register(&DiscoveryNATS{})
	//RegisterDiscovery(&DiscoveryETCD{})
}

func Register(discovery face.IDiscovery) {
	if discovery == nil {
		clog.Fatal("Discovery instance is nil")
		return
	}

	if discovery.Name() == "" {
		clog.Fatalf("Discovery name is empty. %T", discovery)
		return
	}
	discoveryMap[discovery.Name()] = discovery
}
