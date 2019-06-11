package common

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
)

var (
	ConfData *viper.Viper
	confMutex sync.RWMutex
)

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
	loadConf("iris.local.yml")
	ConfData = viper.GetViper()
	loadConf("iris.yml")
	loadConf(fmt.Sprintf("iris.%s.yml",GetEnv()))
}

func init(){
	initConfig()
}
