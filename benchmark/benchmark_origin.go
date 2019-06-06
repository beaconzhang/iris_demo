package main

import (
    "net/http"
    "time"
    "sync"
    "flag"
    "fmt"
    "strconv"
    "io"
    "encoding/json"
)

type ArgumentForm struct{
    Name string `json:name`
    Age int `json:age`
    Count int `json:count`
}

var (
    getCount = 0
    getMutex sync.Mutex
    postCount = 0
    postMutex sync.Mutex
    sleepTime int
)

func SleepTime(st int){
    time.Sleep(time.Duration(st)*time.Millisecond)
}

func getHandler(w http.ResponseWriter, r *http.Request){
    getMutex.Lock()
    getCount += 1
    getMutex.Unlock()
    SleepTime(sleepTime)
    io.WriteString(w,fmt.Sprintf("%s,current visit %d",r.URL.Path,getCount))
}

func postHandler(w http.ResponseWriter, r *http.Request){
    var user ArgumentForm
    postMutex.Lock()
    postCount += 1
    user.Count = postCount
    postMutex.Unlock()
    user.Name = r.PostFormValue("Name")
    user.Age,_ = strconv.Atoi(r.PostFormValue("Age"))
    ret,_ := json.Marshal(user)
    SleepTime(sleepTime)
    w.Write(ret)
}

func init(){
    flag.IntVar(&sleepTime,"s",0,"sleep time")
}

func main(){
    flag.Parse()
    fmt.Printf("sleepTime:%d\n",sleepTime)
    http.HandleFunc("/hello",getHandler)
    http.HandleFunc("/world",postHandler)
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
		fmt.Print("ListenAndServe: ", err)
	}
}
