package main

import (
	"fmt"
	"github.com/juniorh/sample-go-app-1/p1"
//	"sample-go-app-1/p1"	//works if module name in go.mod is the same
)

func main() {
	fmt.Println("Run app1")
	p1.Test()
}
