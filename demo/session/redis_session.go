package main

import (
	"github.com/beaconzhang/iris_demo/common"
	"github.com/beaconzhang/iris_demo/common/session"
	"github.com/beaconzhang/iris_demo/common/storage"
	"github.com/beaconzhang/iris_demo/middleware"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)




func main(){
	confData := common.GetConfData()
	sessionConf := confData.Session
	redisConf := confData.Redis
	golog.Infof("redisConf:%v sessionConf:%v\n",redisConf,sessionConf)
	db := storage.GetRedisClient()
	iris.RegisterOnInterrupt(func(){
		db.Close()
	})
	defer db.Close()

	//sess := sessions.New(sessions.Config{
	//	Cookie: sessionConf.CookieId,
	//	Expires: time.Duration(sessionConf.Expire) * time.Minute,
	//})
	//
	//sess.UseDatabase(db)
	sess := session_manager.New()
	app := iris.New()
	app.Use(middleware.RequestIdMiddlerware)

	app.Get("/",func(ctx iris.Context){
		ctx.Writef("You should navigat to the /set /get /delete /clear /destroy instead")
	})

	app.Get("/set",func(ctx iris.Context){
		s := sess.Start(ctx)
		//set session value
		s.Set("name","iris")
		ctx.Writef("All ok session value of the 'name' is '%s'",s.GetString("name"))
	})

	app.Get("/set/int/{key}/{value}",func(ctx iris.Context){
		key := ctx.Params().GetString("key")
		value := ctx.Params().GetString("value")
		s := sess.IsLogin(ctx)
		if s == nil{
			common.InnerLoggerErrorf(ctx,"no cookie id,user not login")
			s = sess.Start(ctx)
			s.Set(sessionConf.EmployeeInfo.Id,iris.Map{"id":"2726782"})
		}
		employeeId := s.Get(sessionConf.EmployeeInfo.Id).(map[string]interface{})["id"].(string)
		common.InnerLoggerInfof(ctx,"employee id:%s",employeeId)
		s.Set(key,value)
		valueSet := s.Get(key).(string)
		ctx.Writef("All ok session value of the '%s' is '%V'",key,valueSet)
	})

	app.Get("/get/{key}",func(ctx iris.Context){
		key := ctx.Params().Get("key")
		value := sess.Start(ctx).Get(key)
		ctx.Writef("the '%s' on the /set was:'%V'",key,value)
	})

	app.Get("/get",func(ctx iris.Context){
		allSession := sess.Start(ctx).GetAll()
		for k,v := range allSession{
			ctx.Writef("key:%s value %v\n",k,v)
		}
	})

	app.Get("/delete/{key}",func(ctx iris.Context){
		key := ctx.Params().GetString("key")
		sess.Start(ctx).Delete(key)
	})

	app.Get("/clear",func(ctx iris.Context){
		sess.Start(ctx).Clear()
	})

	app.Get("/destroy",func(ctx iris.Context){
		sess.Destroy(ctx)
	})

	app.Get("/update",func(ctx iris.Context){
		if err := sess.ShiftExpiration(ctx); err != nil {
			if sessions.ErrNotFound.Equal(err){
				ctx.StatusCode(iris.StatusNotFound)
			}else if sessions.ErrNotImplemented.Equal(err){
				ctx.StatusCode(iris.StatusNotImplemented)
			}else{
				ctx.StatusCode(iris.StatusNotModified)
			}
			ctx.Writef("%v",err)
			common.InnerLoggerErrorf(ctx,"error:%s",err.Error())
		}
	})
	app.Run(iris.Addr(":8081"),iris.WithoutServerError(iris.ErrServerClosed))

}