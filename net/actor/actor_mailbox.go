package cherryActor

import (
	"time"

	creflect "github.com/po2656233/superplace/extend/reflect"
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
)

type mailbox struct {
	queue                                 // queue
	name    string                        // 邮箱名
	funcMap map[string]*creflect.FuncInfo // 已注册的函数
}

func newMailbox(name string) mailbox {
	return mailbox{
		queue:   newQueue(),
		name:    name,
		funcMap: make(map[string]*creflect.FuncInfo),
	}
}

func (p *mailbox) Register(funcName string, fn interface{}) {
	if funcName == "" || len(funcName) < 1 {
		clog.Errorf("[%s] Func name is empty.", fn)
		return
	}

	funcInfo, err := creflect.GetFuncInfo(fn)
	if err != nil {
		clog.Errorf("funcName = %s, err = %v", funcName, err)
		return
	}

	if _, found := p.funcMap[funcName]; found {
		clog.Errorf("funcName = %s, already exists.", funcName)
		return
	}

	p.funcMap[funcName] = &funcInfo
}

func (p *mailbox) GetFuncInfo(funcName string) (*creflect.FuncInfo, bool) {
	funcInfo, found := p.funcMap[funcName]
	return funcInfo, found
}

func (p *mailbox) Pop() *cfacade.Message {
	v := p.queue.Pop()
	if v == nil {
		return nil
	}

	msg, ok := v.(*cfacade.Message)
	if !ok {
		clog.Warnf("Convert to *Message fail. v = %+v", v)
		return nil
	}

	return msg
}

func (p *mailbox) Push(m *cfacade.Message) {
	if m != nil {
		m.PostTime = time.Now().UnixMilli()
		p.queue.Push(m)
	}
}

func (p *mailbox) onStop() {
	for key := range p.funcMap {
		delete(p.funcMap, key)
	}

	p.queue.Destroy()
}
