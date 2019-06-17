package common

import "github.com/kataras/iris"

func GetXRequestId(ctx iris.Context)string{
	requestIdValue := ctx.GetHeader(constRequestHeader)
	if requestIdValue == ""{
		requestIdValue = ctx.Values().GetString(constRequestHeader)
	}
	return requestIdValue
}
