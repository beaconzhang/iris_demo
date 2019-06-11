// https://blog.csdn.net/wm5920/article/details/77198153
// https://godoc.org/gopkg.in/yaml.v2
package main

import (
    "log"
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

var confYamlFileName = "conf.yml"

type ConfData struct{
    IpPort string `yaml:ipport`
    StartSendTime string `yaml:startsendtime`
    SendMaxCountPerday int `yaml:sendmaxcountperday`
    Devices []Device `yaml:devices`
    WarnFrequncey int
    SendFrequency int
    Other map[string]interface{} `yaml:other`
}

type Device struct{
    DevId string `yaml:devid`
    Nodes []Node `yaml:nodes`
}

type Node struct{
    PkId int `yaml:pkid`
    BkId int `yaml:bkid`
    Index int `yaml:index`
    MinValue int
    MaxValue int
    DataType string
}

func parseYml(filename string) *ConfData{
    data,err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalf("read '%s' error message:%s",filename,err.Error())
        panic(err)
    }
    // https://blog.csdn.net/li_101357/article/details/80209413
    conf := ConfData{}
    yaml.Unmarshal(data,&conf)
    return &conf
}

func main(){
    conf := parseYml(confYamlFileName)
    fmt.Printf("%V\n",*conf)
    marshalConf,_ := yaml.Marshal(conf)
    fmt.Printf("marshal:\n%s\n",marshalConf)
}
