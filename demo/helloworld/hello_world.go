package main

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/middleware/recover"
    "github.com/kataras/iris/middleware/logger"
    "github.com/beaconzhang/iris_demo/common"
    "github.com/beaconzhang/iris_demo/middleware"
)

//func handlerLoggerMiddlerware(ctx iris.Context){
//    requestId := ctx.GetHeader("x_request_id")
//    ctx.Application().Logger().Prefix = []byte("["+requestId+"]")
//    ctx.Next()
//}


func main(){
    app := iris.New()
    app.Logger().SetLevel("debug")
    app.Use(middleware.RequestIdMiddlerware)
    app.Use(recover.New())
    loggerConfig := logger.DefaultConfig()
    loggerConfig.MessageHeaderKeys = make([]string,1)
    loggerConfig.MessageHeaderKeys = append(loggerConfig.MessageHeaderKeys,"x_request_id")
    loggerConfig.MessageContextKeys = append(loggerConfig.MessageContextKeys,"x_request_id")
    logger := logger.New(loggerConfig)
    app.Use(logger)
    //app.Use(handlerLoggerMiddlerware)

    app.Handle("GET","/",func(ctx iris.Context){
        common.InnerLoggerInfof(ctx,"get hello")
        ctx.HTML("<h1>welcome</h1>")
        common.InnerLoggerErrorf(ctx,"done hello")
    })

    app.Handle("GET","/ping",func(ctx iris.Context){
        ctx.Application().Logger().Errorf("get hello")
        ctx.WriteString("pong")
    })
    app.Get("/hello",func(ctx iris.Context){
        ctx.JSON(iris.Map{"message":"hello iris"})
    })
    app.Run(iris.Addr(":8081"),iris.WithoutServerError(iris.ErrServerClosed))
}
