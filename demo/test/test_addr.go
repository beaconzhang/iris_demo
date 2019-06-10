package main
import (
    "fmt"
)

type Astruct struct{
    a int
    b string
}

func NewAstruct() *Astruct{
    return &Astruct{
        a:3,
    }
}

func main(){
    a1 := NewAstruct()
    a2 := NewAstruct()
    fmt.Printf("before change: a1 addr:%p value:%v a2 addr:%p value:%v\n",a1,a1,a2,a2)
}
