package auth

import (
	"github.com/beaconzhang/iris_demo/common"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var (
	pathLogin = mvc.Response{Path: "/auth/user/login"}
	pathLogout = mvc.Response{Path: "/auth/user/logout"}
)

type AuthUserController struct{

}

type templateLoginField struct{
	Title string
	GithubUrl string
}

func (ac *AuthUserController)BeginRequest(ctx iris.Context){}

func (ac *AuthUserController)EndRequest(ctx iris.Context){}
func (c *AuthUserController) BeforeActivation(b mvc.BeforeActivation) {

	b.Dependencies().Add(func(ctx iris.Context) iris.Context { return ctx })
}

func (ac *AuthUserController)GetUserLogin(ctx iris.Context) mvc.Result{
	common.InnerLoggerInfof(ctx,"/auth/user/login in")
	return  mvc.View{
		Name:  "user/login.html",
		Data: templateLoginField{Title:"User login",GithubUrl:"/auth/user/github/login"},
	}
}
