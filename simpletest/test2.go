package main

import (
	"fmt"
	"test"
)

func main() {
	var a *test.Test1 = new(test.Test1)
	a.Name = "zzzz"
	test.Test2()

	fmt.Print("Test1.Name:", a.Name)
}
