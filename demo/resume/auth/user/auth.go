package auth

import (
	"github.com/beaconzhang/iris_demo/common"
	"github.com/beaconzhang/iris_demo/common/thirtypart_auth/github"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"net/url"
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
	next := ctx.FormValueDefault("next","/")
	redirectUrl :="/auth/github/login?"+ "next="+url.QueryEscape(next)
	common.InnerLoggerInfof(ctx,"redirectUrl:%s",redirectUrl)
	return  mvc.View{
		Name:  "user/login.html",
		Data: templateLoginField{Title:"User login",GithubUrl:redirectUrl},
	}
}

func (ac *AuthUserController)GetGithubLogin(ctx iris.Context) {
	githubAuth := &github.GithubAuth{}
	githubAuth.GetIdentify(ctx)
}

func (ac *AuthUserController)GetGithubCallback(ctx iris.Context){
	githubAuth := &github.GithubAuth{}
	githubAuth.GetUserInfo(ctx)
}
