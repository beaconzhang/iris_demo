package main

import (
	"github.com/kataras/iris"
    "time"
    "sync"
    "flag"
    "fmt"
)

type ArgumentForm struct{
    Name string `json:name`
    Age int `json:age`
    Count int `json:count`
}

var (
    getCount = 0
    getMutex sync.Mutex
    postCount = 0
    postMutex sync.Mutex
    sleepTime int
)

func SleepTime(st int){
    time.Sleep(time.Duration(st)*time.Millisecond)
}

func getHandler(ctx iris.Context){
    getMutex.Lock()
    getCount += 1
    getMutex.Unlock()
    SleepTime(sleepTime)
    ctx.Writef("%s,current visit %d",ctx.Path(),getCount)
}

func postHandler(ctx iris.Context){
    var user ArgumentForm
    postMutex.Lock()
    postCount += 1
    user.Count = postCount
    postMutex.Unlock()
    err := ctx.ReadForm(&user)
    if err != nil {
		ctx.Writef("error","create user,read or parse form failed"+err.Error())
        ctx.StatusCode(iris.StatusInternalServerError)
        return
    }
    SleepTime(sleepTime)
    ctx.JSON(user)

}

func init(){
    flag.IntVar(&sleepTime,"s",0,"sleep time")
}

func main(){
    flag.Parse()
    fmt.Printf("sleepTime:%d\n",sleepTime)
    app := iris.New()
    app.Get("/hello",getHandler)
    app.Post("/world",postHandler)
    app.Run(iris.Addr(":8081"),iris.WithCharset("utf-8"))
}
