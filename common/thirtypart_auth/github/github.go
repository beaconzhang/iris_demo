package github

import (
	"encoding/json"
	"github.com/beaconzhang/iris_demo/common"
	"github.com/kataras/iris"
	"io/ioutil"
	"net/http"
	"strings"
	"errors"
)

type GithubAuth struct{
}


func (ga *GithubAuth)GetIdentify(ctx iris.Context) {
	reqUrl := ga.Redirect(ctx,"")
	ctx.Redirect(reqUrl,iris.StatusFound)
}

func (ga *GithubAuth) Redirect(ctx iris.Context,next string) string{
	if next == "" {
		next = ctx.FormValueDefault("next", "/")
	}

	state := map[string]string {
		"id":common.GetXRequestId(ctx),
		"next": next,
	}
	stateStr,_ := json.Marshal(state)

	confData := common.GetConfData()
	//redisClient := storage.GetOriginRedisClient()
	//redisClient.Set(state,"1",time.Duration(3*60))
	reqUrl := confData.Auth.Github.GithubIdentityUrl+"?" + "next=" + next +
		"&client_id="+confData.Auth.Github.GithubClientId+"&state="+string(stateStr)
	common.InnerLoggerInfof(ctx,"redirct to github url:%s stateStr:%s, next:%s",
		reqUrl,stateStr,next)
	return reqUrl
}

func (ga *GithubAuth)GetToken(ctx iris.Context,code string,state string)(string,error){
	confData := common.GetConfData()
	githubConf := &confData.Auth.Github
	reqUrl := githubConf.GithubTokenUrl + "?client_id=" + githubConf.GithubClientId + "&client_secret=" +
		githubConf.GithubClientSecret + "&code=" + code + "&state=" + state
	reqClient := &http.Client{}
	reqPost, err := http.NewRequest("POST",reqUrl,strings.NewReader(""))
	if err != nil{
		return "newRequest error",err
	}
	reqPost.Header.Set("content-type","application/x-www-form-urlencoded")
	reqPost.Header.Set("accept","application/json")
	reqPost.Header.Set("x_request_id",common.GetXRequestId(ctx))
	resp,err:=reqClient.Do(reqPost)
	if err != nil{
		return "post http request error",err
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return "read body error",err
	}
	respJson := map[string]interface{}{}
	err = json.Unmarshal(body,&respJson)
	if err !=nil{
		return "unmarshal "+string(body) +" error",err
	}
	token,okToken := respJson["access_token"]
	scope,okScope := respJson["scope"]
	tokenType,okTokenType := respJson["token_type"]
	if !okToken || !okTokenType || !okScope{
		return "get field from "+ string(body)+ " error",errors.New("parse error")
	}
	tokenStr := token.(string)
	scopeStr := scope.(string)
	tokenTypeStr := tokenType.(string)
	common.InnerLoggerInfof(ctx,"reqUrl:%s token:%s scope:%s tokenType:%s",tokenStr,scopeStr,tokenTypeStr)
	return tokenStr,nil
}

func (ga *GithubAuth)getUserInfo(ctx iris.Context,token string)(string,error){
	confData := common.GetConfData()
	githubConf := &confData.Auth.Github
	reqUrl := githubConf.GithubApiUrl
	reqClient := &http.Client{}
	reqGet,err := http.NewRequest("GET",reqUrl,strings.NewReader(""))
	if err != nil {
		return "error init request:" + reqUrl,err
	}
	reqGet.Header.Set("Authorization","token "+token)
	reqGet.Header.Set("x_request_id",common.GetXRequestId(ctx))
	resp,err := reqClient.Do(reqGet)
	if err != nil {
		return "error get request:" + reqUrl,err
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "error read body ",err
	}
	return string(body),nil
}

func (ga *GithubAuth)GetUserInfo(ctx iris.Context){
	githubCode := ctx.FormValue("code")
	githubState := ctx.FormValue("state")

	var stateJson  map[string]interface{}
	err := json.Unmarshal([]byte(githubState),&stateJson)
	if err != nil{
		ga.Redirect(ctx,"")
		return
	}
	next := stateJson["next"].(string)
	id := stateJson["id"].(string)

	common.InnerLoggerInfof(ctx,"code:%s state:%s request_path:%snext:%s id:%s",
		githubCode,githubState,ctx.Request().RequestURI,next,id)

	token,err := ga.GetToken(ctx,githubCode,githubState)
	if err != nil {
		common.InnerLoggerErrorf(ctx,"get token error,%s error:%s",token,err.Error())
		ga.Redirect(ctx,next)
		return
	}

	body,err := ga.getUserInfo(ctx,token)
	if err != nil{
		common.InnerLoggerErrorf(ctx,"get user info error,%s error:%s",body,err.Error())
		ga.Redirect(ctx,next)
		return
	}
	common.InnerLoggerInfof(ctx,"get user info:%s",body)
}
