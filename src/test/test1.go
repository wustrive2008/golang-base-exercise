package test

import (
	"fmt"
)

type Test1 struct {
	Name string
}

func (a *Test1) Test() string {
	return a.Name
}

func Test2() {
	fmt.Print("test.Test2() invoked...")
}
