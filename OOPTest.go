package main

import (
	"fmt"
)

type Integer int

//传递引用
func (a *Integer) Add(b Integer) {
	*a += b
}

//传递值
func (a Integer) Add2(b Integer) {
	a += b
}
func main() {
	var a Integer = 1
	a.Add2(2)
	fmt.Println("a=", a)
}
