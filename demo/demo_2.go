package main

import (
    "github.com/kataras/iris"
    "strconv"
)


func main(){
    app := iris.New()
    app.Get("/username/{name}",func(ctx iris.Context){
        ctx.Writef("hello:%s",ctx.Params().Get("name"))
    })
    app.Macros().Get("int").RegisterFunc("min",func(minValue int)func(string) bool{
        return func(paramValue string)bool{
            n,err := strconv.Atoi(paramValue)
            if err != nil{
                app.Logger().Fatal("erro:%s",err.Error())
                return false
            }
            return minValue <= n
        }
    })
    app.Get("/profile/{id:int min(3)}",func(ctx iris.Context){
        id,_ := ctx.Params().GetInt("id")
        ctx.Writef("hello profile id:%d",id)
    })

    app.Get("/profile/{name:int min(3)}/friends/{friendid:int min(1) else 504}",
        func(ctx iris.Context){
            id,_ := ctx.Params().GetInt("name")
            friendId,_ := ctx.Params().GetInt("friendid")
            ctx.Writef("hello profile id:%d looking for friend:%d",id,friendId)
        })

    app.Get("/game/{name:alphabetical}/level/{level:int}",func(ctx iris.Context){
        ctx.Writef("name:%s | level:%s",ctx.Params().Get("name"),ctx.Params().Get("level"))
    })

    app.Get("/lowercase/{name:string regexp(^[a-z]+)}",func(ctx iris.Context){
        ctx.Writef("name should be only lowercase,otherwise this handler will never execute:%s",ctx.Params().Get("name"))

    })

    app.Get("/single_file/{myfile:file}",func(ctx iris.Context){
        ctx.Writef("file type validate if parameter value has a form of file:%s",ctx.Params().Get("myfile"))
    })
    app.Get("/myfile/{directory:path}",func(ctx iris.Context){
        ctx.Writef("path type accepts any number of path segment,path after /myfile/ is:%s",ctx.Params().Get("directory"))
    })
    app.Run(iris.Addr(":8081"))

}
