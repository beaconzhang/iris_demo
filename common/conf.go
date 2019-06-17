package common

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
)

var (
	confDataInterface *viper.Viper
	confMutex sync.Mutex
	confDataParse *ConfData
)

type ConfData struct{
	Redis redisConf
	Session sessionConf
	Auth authConf
	WhitelistUrl []string
}

type redisConf struct{
	Network string
	Host string
	Port string
	Passwd string
	Database string
	Maxidle int
	Maxactive int
	Idletimeout int
	Prefix string
}

type sessionConf struct{
	CookieId string
	Expire int
	EmployeeInfo userRecord
}

type userRecord struct{
	Prefix string
}

type authConf struct{
	LoginUrl string
	Github githubConf
}

type githubConf struct{
	GithubLoginUrl string
	GithubCallBackUrl string
	GithubIdentityUrl string
	GithubTokenUrl string
	GithubApiUrl string
	GithubClientId string
	GithubClientSecret string
	GithubScope string
}



func GetConfData() *ConfData{
	if confDataParse != nil{
		return confDataParse
	}
	confMutex.Lock()
	defer confMutex.Unlock()
	if confDataParse != nil{
		return confDataParse
	}
	confDataParse = &ConfData{}
	err := confDataInterface.Unmarshal(confDataParse)
	golog.Errorf("read conf data error:%s",err.Error())
	return confDataParse
}

func loadConf(filename string){
	rootDir := GetRootDir()
	localConfFilePath := filepath.Join(rootDir,"conf",filename)
	if _,ok := os.Stat(localConfFilePath); os.IsNotExist(ok){
		return
	}
	localConfByte,_ := os.Open(localConfFilePath)
	defer localConfByte.Close()
	confMutex.Lock()
	defer confMutex.Unlock()
	err :=viper.MergeConfig(localConfByte)
	if err != nil {
		fmt.Printf("read filepath:%s erro:%s",localConfFilePath,err.Error())
	}
	fmt.Printf("filepat:%s\n",localConfFilePath)
}


func initConfig(){
	viper.Reset()
	viper.SetConfigType("yaml")
	confDataInterface = viper.GetViper()
	loadConf("iris.yml")
	loadConf(fmt.Sprintf("iris.%s.yml",GetEnv()))
	loadConf("iris.local.yml")
}

func init(){
	initConfig()
	confData := GetConfData()
	golog.Infof("read conf data:%v",confData)
}

