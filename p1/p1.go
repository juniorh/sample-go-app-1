package p1

import "fmt"

func Test() {
    fmt.Println("Call from package p1.Test()")
}

func Sapi() {
    fmt.Println("Call from package p1.Sapi()")
}

func cacing() {
    fmt.Println("This can't be called from outside package since since function name start with lowercase")
}
