package cherryActor

import (
	"strings"
	"sync"

	cfacade "github.com/po2656233/superplace/facade"
)

type actorChild struct {
	thisActor   *Actor
	childActors *sync.Map // key:childActorId, value:*actor
}

func newChild(thisActor *Actor) actorChild {
	return actorChild{
		thisActor:   thisActor,
		childActors: &sync.Map{},
	}
}

func (p *actorChild) onStop() {
	p.childActors.Range(func(key, value any) bool {
		if childActor, ok := value.(*Actor); ok {
			childActor.Exit()
		}
		return true
	})

	//p.childActors = nil
	p.thisActor = nil
}

func (p *actorChild) Create(childID string, handler cfacade.IActorHandler) (cfacade.IActor, error) {
	if p.thisActor.path.IsChild() {
		return nil, ErrForbiddenCreateChildActor
	}

	if strings.TrimSpace(childID) == "" {
		return nil, ErrActorIDIsNil
	}

	if thisActor, ok := p.Get(childID); ok {
		return thisActor, nil
	}

	childActor, err := newActor(p.thisActor.ActorID(), childID, handler, p.thisActor.system)
	if err != nil {
		return nil, err
	}

	p.childActors.Store(childID, &childActor)
	go childActor.run()

	return &childActor, nil
}

func (p *actorChild) Get(childID string) (cfacade.IActor, bool) {
	return p.GetActor(childID)
}

func (p *actorChild) GetActor(childID string) (*Actor, bool) {
	if actorValue, ok := p.childActors.Load(childID); ok {
		actor, found := actorValue.(*Actor)
		return actor, found
	}

	return nil, false
}

func (p *actorChild) Remove(childID string) {
	p.childActors.Delete(childID)
}

func (p *actorChild) Each(fn func(cfacade.IActor)) {
	p.childActors.Range(func(key, value any) bool {
		if actor, found := value.(*Actor); found {
			fn(actor)
		}
		return true
	})
}
