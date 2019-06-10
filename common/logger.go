package common

import (
    "github.com/kataras/iris"
)


func innerLogger(ctx iris.Context,loggerHandler func(string, ...interface{}),format string,args ...interface{}){
    args = append(args,ctx.GetHeader("x_request_id"))
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
