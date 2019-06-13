package main

import (
	"github.com/beaconzhang/iris_demo/common"
	"github.com/beaconzhang/iris_demo/demo/resume/mvc"
	"github.com/beaconzhang/iris_demo/middleware"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/mvc"
)

func main(){
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	loggerConfig := logger.DefaultConfig()
	loggerConfig.MessageHeaderKeys = make([]string,1)
	loggerConfig.MessageHeaderKeys = append(loggerConfig.MessageHeaderKeys,"x_request_id")
	loggerConfig.MessageContextKeys = append(loggerConfig.MessageContextKeys,"x_request_id")
	logger := logger.New(loggerConfig)
	app.Use(logger)
	app.Use(middleware.RequestIdMiddlerware)
	app.Use(middleware.SessionMiddlerware)

	app.RegisterView(iris.HTML( "./demo/resume/views", ".html").Layout("share/layout.html"))
	app.StaticWeb("/public", "./demo/resume/public")

	mvc.Configure(app,resume_mvc.ConfigureAuthMvc)
	mvc.Configure(app,resume_mvc.ConfiguerHelloMvc)
	app.OnErrorCode(400, func(ctx iris.Context) {
		println("here") // this will be executed on http(s)://$host/error request path.
		reason := ctx.Values().Get("reason")
		ctx.WriteString(reason.(string))
	})


	app.Get("/hei/ha",func(ctx iris.Context){
		ctx.ViewData("Title","hello")
		ctx.ViewData("Firstname","github")
		common.InnerLoggerInfof(ctx,"login")
		ctx.View("user/test.html")
	})
	//app.Get("/hello",func(ctx iris.Context){
	//	ctx.Writef("hei ha")
	//	common.InnerLoggerInfof(ctx,"login")
	//})
	app.Run(iris.Addr(":8081"))
}
