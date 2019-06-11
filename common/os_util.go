package common
//package main

import (
    "net"
    "fmt"
    "os"
    "strings"
)

func GetMac(name string) (string,string) {
    ifas,err := net.Interfaces()
    if err != nil {
        return "", ""
    }
    for _,ifa := range ifas {
        fmt.Printf("%s\n",ifa.Name)
        var mac string
        var ip string
        if ifa.Name == name {
            mac = ifa.HardwareAddr.String()
            mac = strings.Replace(mac,":","",-1)
            addrs, err := ifa.Addrs()
            if err != nil {
                ip = ""
            }else{
                for _,addr := range addrs{
                    if ipv4,ok := addr.(*net.IPNet);ok && !ipv4.IP.IsLoopback() {
                        ipv4Byte := ipv4.IP.To4()
                        if  ipv4Byte != nil{
                            ip = fmt.Sprintf("%02x%02x%02x%02x",ipv4Byte[0],ipv4Byte[1],ipv4Byte[2],ipv4Byte[3])
                            break
                        }
                    }
                }
            }
            return mac, ip
        }
    }
    return "",""
}

func GetPid()string{
    pid := os.Getpid()
    //fmt.Printf("pid:%d\n",pid)
    return fmt.Sprintf("%4x",pid)
}

//func main(){
//   mac,ip := GetMac("en0")
//   fmt.Printf("mac:%s ip:%s\n",mac,ip)
//   //hostname,_ := os.Hostname()
//   //fmt.Printf("hostname:%s\n",hostname)
//   fmt.Printf("pid:%s\n",GetPid())
//}
