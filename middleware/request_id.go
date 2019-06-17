// https://github.com/go-chi/chi/blob/master/middleware/request_id.go

package middleware

import (
    "encoding/base64"
    "github.com/beaconzhang/iris_demo/common"
    "github.com/kataras/iris"
    "strings"
    "crypto/rand"
    "fmt"
    "sync/atomic"
    "time"
)

func getRandom() string{
    var buf [12]byte
    var b64 string
    for len(b64) < 5 {
        rand.Read(buf[:])
        b64 = base64.StdEncoding.EncodeToString(buf[:])
        b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
    }
    return b64[:5]
}
var prefix string
var requestId uint64
var constPrefix string
var hashSalt string
var constRequestHeader = "x_request_id"

func init(){
    var ethnetName = "en0"
    prefix = "hello"
    hashSalt = "world"
    mac,ip := common.GetMac(ethnetName)
    pid := common.GetPid()
    constPrefix  = fmt.Sprintf("%s%s%s%s",prefix,mac,ip,pid)
    requestId = 0
}

func getRequestId() string{
    myid := atomic.AddUint64(&requestId, 1)
    myid = myid & 0xffffffff
    return fmt.Sprintf("%s%s%08x%08x",constPrefix,getRandom(),(time.Now().Unix())&0xffffffff,myid)
}

func RequestIdMiddlerware(ctx iris.Context){
   requestIdValue := ctx.GetHeader(constRequestHeader)
   if requestIdValue == ""{
       requestIdValue = getRequestId()
       ctx.Values().Set(constRequestHeader,requestIdValue)
   }
    ctx.Header(constRequestHeader,requestIdValue)
   //ctx.Application().Logger().Prefix = []byte("["+requestIdValue+"]")
   ctx.Next()
}




