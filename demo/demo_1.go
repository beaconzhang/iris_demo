// https://studyiris.com/example/iris.html
package main

import (
	"github.com/kataras/iris"
	"fmt"
)

type User struct{
	Username string `json:"username"`
	Firstname string `json:firstname`
	Lastname string `json:lastname`
	City string `json:city`
	Age int `json:age`
}

func (user *User) String()string{
	return fmt.Sprintf("username:%s firstname:%s lastname:%s city:%s age:%d",
		user.Username,user.Firstname,user.Lastname,user.City,user.Age)
}



func main(){
	app := iris.New()
	app.RegisterView(iris.HTML("./web/views",".html").Reload(true))
	app.OnErrorCode(iris.StatusInternalServerError,func(ctx iris.Context){
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error:%s",errMessage)
			return
		}
		ctx.Writef("(Unexpected) internal server error")
	})
	app.Use(func(ctx iris.Context){
		ctx.Application().Logger().Infof("Begin request for path:%s",ctx.Path())
		ctx.Next()
	})
	app.Done(func(ctx iris.Context){})
	app.Post("/decode",func(ctx iris.Context){
		var user User
		ctx.ReadJSON(&user)
		ctx.Writef("%s",user.String())
	})
	app.Get("/encode",func(ctx iris.Context){
		doe := User{
			Username: "Johndoe",
			Firstname: "john",
			Lastname:"doe",
			City:"Neither FBI knows!!!",
			Age: 25,
		}
		ctx.JSON(doe)
	})

	app.Get("/profile/{username:string}",profileByUsername)
	userRouters := app.Party("/users",logThisMiddleware)
	{
		userRouters.Get("/{id:int min(1)}",getUserByID)
		userRouters.Post("/create",createUser)
	}

	app.Run(iris.Addr(":8081"),iris.WithCharset("utf-8"))

}

func logThisMiddleware(ctx iris.Context){
	ctx.Application().Logger().Infof("Path:%s | IP:%s",ctx.Path(),ctx.RemoteAddr())
	ctx.Next()
}

func profileByUsername(ctx iris.Context){
	username := ctx.Params().Get("username")
	ctx.ViewData("Username",username)
	ctx.View("profile.html")
}

func getUserByID(ctx iris.Context){
	userId := ctx.Params().Get("id")
	user := User{Username:"username" + userId}
	ctx.XML(user)
}

func createUser(ctx iris.Context){
	var user User
	err := ctx.ReadForm(&user)
	if err != nil{
		ctx.Values().Set("error","create user,read or parse form failed"+err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	ctx.ViewData("",user)
	ctx.View("create_verification.html")

}
