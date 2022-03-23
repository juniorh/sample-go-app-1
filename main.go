package main

import "fmt"
import "github.com/juniorh/sample-go-app-1/p1"  // PATH should match with module name in go.mod
import "github.com/juniorh/sample-go-app-1/p2"
//import "app1/p1"
//import "app1/p2"

func main() {
	fmt.Println("Good Morning!")
	f1()
	p1.Test()
	p1.Sapi()
	//p1.cacing()
	p2.Bebek()
	fmt.Println("End Script!")
}
