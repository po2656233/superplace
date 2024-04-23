package cherryActor

import (
	"time"

	creflect "github.com/po2656233/superplace/extend/reflect"
	face "github.com/po2656233/superplace/facade"
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

// Register 第一个为函数指针 第二个指定函数名(可省略)
func (p *mailbox) Register(fn interface{}) bool {
	funcInfo, err := creflect.GetFuncInfo(fn)
	if err != nil {
		clog.Errorf("funcName = %s, err = %v", funcInfo.Name, err)
		return false
	}
	funcName := funcInfo.Name
	if _, found := p.funcMap[funcName]; found {
		clog.Errorf("funcName = %s, already exists.", funcName)
		return false
	}
	clog.Infof("MailBox Register: [%s] ok", funcName)
	p.funcMap[funcName] = &funcInfo
	return true
}

func (p *mailbox) GetFuncInfo(funcName string) (*creflect.FuncInfo, bool) {
	funcInfo, found := p.funcMap[funcName]
	return funcInfo, found
}

func (p *mailbox) Pop() *face.Message {
	v := p.queue.Pop()
	if v == nil {
		return nil
	}

	msg, ok := v.(*face.Message)
	if !ok {
		clog.Warnf("Convert to *Message fail. v = %+v", v)
		return nil
	}

	return msg
}

func (p *mailbox) Push(m *face.Message) {
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
