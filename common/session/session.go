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
		return nil
	}
	sess := s.Start(ctx)
	employeInfo := sess.Get(sessionConf.EmployeeInfo)
	if employeInfo == nil {
		return nil
	}
	sess.SetFlash(sessionConf.EmployeeInfo,employeInfo)
	return sess
}
