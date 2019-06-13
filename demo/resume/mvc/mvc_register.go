package resume_mvc

import (
	"github.com/beaconzhang/iris_demo/demo/resume/auth/user"
	"github.com/beaconzhang/iris_demo/demo/resume/hello"
	"github.com/kataras/iris/mvc"
)

func ConfigureAuthMvc(app *mvc.Application){
	userApp := app.Party("/auth")
	userApp.Handle(new(auth.AuthUserController))
}

func ConfiguerHelloMvc(app *mvc.Application){
	userApp := app.Party("/hello")
	userApp.Handle(new(hello.HelloController))
}