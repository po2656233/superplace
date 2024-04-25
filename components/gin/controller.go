package superGin

import (
	"github.com/gin-gonic/gin"
	face "github.com/po2656233/superplace/facade"
)

type GinHandlerFunc func(ctx *Context)

type IController interface {
	PreInit(app face.IApplication, engine *gin.Engine)
	Init()
	Stop()
}

func BindHandlers(handlers []GinHandlerFunc) []gin.HandlerFunc {
	var list []gin.HandlerFunc
	for _, handler := range handlers {
		list = append(list, BindHandler(handler))
	}
	return list
}

func BindHandler(handler func(ctx *Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.Context = c
		handler(context)
	}
}

type BaseController struct {
	App    face.IApplication
	Engine *gin.Engine
}

func (b *BaseController) PreInit(app face.IApplication, engine *gin.Engine) {
	b.App = app
	b.Engine = engine
}

func (b *BaseController) Init() {

}

func (b *BaseController) Stop() {

}

func (b *BaseController) Group(relativePath string, handlers ...GinHandlerFunc) *Group {
	group := &Group{
		RouterGroup: b.Engine.Group(relativePath, BindHandlers(handlers)...),
	}
	return group
}

func (b *BaseController) Any(relativePath string, handlers ...GinHandlerFunc) {
	b.Engine.Any(relativePath, BindHandlers(handlers)...)
}

func (b *BaseController) GET(relativePath string, handlers ...GinHandlerFunc) {
	b.Engine.GET(relativePath, BindHandlers(handlers)...)
}

func (b *BaseController) POST(relativePath string, handlers ...GinHandlerFunc) {
	b.Engine.POST(relativePath, BindHandlers(handlers)...)
}
