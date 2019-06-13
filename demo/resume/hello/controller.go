package hello

import (
	"github.com/beaconzhang/iris_demo/common"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type HelloController struct{

}

func (ac *HelloController)BeginRequest(ctx iris.Context){}

func (ac *HelloController)EndRequest(ctx iris.Context){}
func (c *HelloController) BeforeActivation(b mvc.BeforeActivation) {

	b.Dependencies().Add(func(ctx iris.Context) iris.Context { return ctx })
}

func (ac *HelloController)GetWorld(ctx iris.Context) mvc.Result{
	common.InnerLoggerInfof(ctx,"/hello/world")
	return mvc.Response{
		Text: "welcome to happy island",
	}
}
