package main

import (
	"fmt"
)

var a string
var c = make(chan int, 2)

func f() {
	a = "hello world"
	<-c
}

func main() {
	go f()
	c <- 0
	fmt.Println(a)
}
