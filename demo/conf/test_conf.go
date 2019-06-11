package main

import (
	"fmt"
	"github.com/beaconzhang/iris_demo/common"
)

type Conf struct{
	Hacker bool
	Name string
	Hobbies []string
}


func main(){
	fmt.Printf("%V\n",common.ConfData)
	fmt.Printf("%s\n",common.ConfData.GetStringSlice("hobbies"))
	c := Conf{}
	common.ConfData.Unmarshal(&c)
	fmt.Printf("%s\n",c)
}
