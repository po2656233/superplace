package simple

import (
	cerr "github.com/po2656233/superplace/logger/error"
	"sync"

	face "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
)

var (
	lock        = &sync.RWMutex{}
	sidAgentMap = make(map[face.SID]*Agent)   // sid -> Agent
	uidMap      = make(map[face.UID]face.SID) // uid -> sid
)

func BindSID(agent *Agent) {
	lock.Lock()
	defer lock.Unlock()

	sidAgentMap[agent.SID()] = agent
}

func BindUID(sid face.SID, uid face.UID) error {
	if sid == "" {
		return cerr.Errorf("[sid = %s] less than 1.", sid)
	}

	if uid < 1 {
		return cerr.Errorf("[uid = %d] less than 1.", uid)
	}

	lock.Lock()
	defer lock.Unlock()

	agent, found := sidAgentMap[sid]
	if !found {
		return cerr.Errorf("[sid = %s] does not exist.", sid)
	}

	if agent.UID() > 0 && agent.UID() == uid {
		return cerr.Errorf("[uid = %d] has already bound.", agent.UID())
	}

	agent.session.Uid = uid
	uidMap[uid] = sid

	return nil
}

func Unbind(sid face.SID) {
	lock.Lock()
	defer lock.Unlock()

	agent, found := sidAgentMap[sid]
	if !found {
		return
	}

	delete(sidAgentMap, sid)
	delete(uidMap, agent.UID())

	sidCount := len(sidAgentMap)
	uidCount := len(uidMap)
	if sidCount == 0 || uidCount == 0 {
		clog.Infof("Unbind agent sid = %s, sidCount = %d, uidCount = %d", sid, sidCount, uidCount)
	}
}

func GetAgent(sid face.SID) (*Agent, bool) {
	lock.Lock()
	defer lock.Unlock()

	agent, found := sidAgentMap[sid]
	return agent, found
}

func GetAgentWithUID(uid face.UID) (*Agent, bool) {
	if uid < 1 {
		return nil, false
	}

	lock.Lock()
	defer lock.Unlock()

	sid, found := uidMap[uid]
	if !found {
		return nil, false
	}

	agent, found := sidAgentMap[sid]
	return agent, found
}

func ForeachAgent(fn func(a *Agent)) {
	for _, agent := range sidAgentMap {
		fn(agent)
	}
}

func Count() int {
	lock.RLock()
	defer lock.RUnlock()

	return len(sidAgentMap)
}
