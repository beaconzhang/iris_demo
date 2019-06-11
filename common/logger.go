package common

import (
    "github.com/kataras/iris"
)

const (
    constRequestHeader = "x_request_id"
)

func innerLogger(ctx iris.Context,loggerHandler func(string, ...interface{}),format string,args ...interface{}){
    requestIdValue := ctx.GetHeader(constRequestHeader)
    if requestIdValue == ""{
        requestIdValue = ctx.Values().GetString(constRequestHeader)
    }
    args = append(args,requestIdValue)
    loggerHandler(format+" [%s]",args...)
}

func InnerLoggerDebugf(ctx iris.Context, format string,args ...interface{}){
    innerLogger(ctx,ctx.Application().Logger().Debugf,format,args...)
}

func InnerLoggerInfof(ctx iris.Context, format string,args ...interface{}){
    innerLogger(ctx,ctx.Application().Logger().Infof,format,args...)
}


func InnerLoggerWarnf(ctx iris.Context, format string,args ...interface{}){
    innerLogger(ctx,ctx.Application().Logger().Warnf,format,args...)
}

func InnerLoggerErrorf(ctx iris.Context, format string,args ...interface{}){
    innerLogger(ctx,ctx.Application().Logger().Errorf,format,args...)
}

func InnerLoggerFatalf(ctx iris.Context, format string,args ...interface{}){
    innerLogger(ctx,ctx.Application().Logger().Fatalf,format,args...)
}
