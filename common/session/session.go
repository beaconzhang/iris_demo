package session_manager

import (
	"github.com/beaconzhang/iris_demo/common"
	"github.com/beaconzhang/iris_demo/common/storage"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"time"
)

type Sessions struct{
	sessions.Sessions
}

type SessionUserInfo struct{
	Id string
}

func New() *Sessions{
	sessionConf := common.GetConfData().Session
	redisClient := storage.GetRedisClient()
	sess := Sessions{
		 *sessions.New(sessions.Config{
			Cookie: sessionConf.CookieId,
			Expires: time.Duration(sessionConf.Expire) * time.Minute,
		}),
	}
	sess.UseDatabase(redisClient)
	return &sess
}

func (s *Sessions)IsLogin(ctx iris.Context) *sessions.Session{
	sessionConf := common.GetConfData().Session
	cookieValue := sessions.GetCookie(ctx, sessionConf.CookieId)
	if cookieValue == ""{
		common.InnerLoggerInfof(ctx,"[None] login")
		return nil
	}
	sess := s.Start(ctx)
	employeInfo := sess.Get(sessionConf.EmployeeInfo.Prefix)
	if employeInfo == nil {
		s.Destroy(ctx)
		common.InnerLoggerInfof(ctx,"[session expire] login")
		return nil
	}
	common.InnerLoggerInfof(ctx,"[%s] login",employeInfo.(SessionUserInfo).Id)
	sess.SetFlash(sessionConf.EmployeeInfo.Prefix,employeInfo)
	ctx.Values().Set("iris_session",sess)
	return sess
}
