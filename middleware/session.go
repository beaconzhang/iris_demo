// 用户是否登录检测，如果没有登录将会重定向到登录页面
// 如果登录，将用户信息记录于redis中，并生成相应cookie
package middleware

import (
	"github.com/beaconzhang/iris_demo/common"
	"github.com/beaconzhang/iris_demo/common/session"
	"github.com/kataras/iris"
	"net/url"
)

var (
	mapWhitelistUrl map[string]bool
)

func init(){
	confData := common.GetConfData()
	mapWhitelistUrl = make(map[string]bool)
	for _,item := range confData.WhitelistUrl{
		mapWhitelistUrl[item] = true
	}
}

func SessionMiddlerware(ctx iris.Context){
	sessesion := session_manager.New()
	sess := sessesion.IsLogin(ctx)
	if sess == nil  {
		param := ctx.Request().RequestURI
		path := ctx.Path()
		if _,ok := mapWhitelistUrl[path];!ok {
			if ctx.IsAjax() {
				ctx.JSON(iris.Map{"sucess": false, "next": param})
				return
			}
			confData := common.GetConfData()
			encodedPath := url.QueryEscape(param)
			ctx.Redirect(confData.Auth.LoginUrl + "?next=" + encodedPath,iris.StatusFound)
			return
		}
		common.InnerLoggerInfof(ctx,"[%s] in whitelist",param)
	}
	ctx.Next()
}



