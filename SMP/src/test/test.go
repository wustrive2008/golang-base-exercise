package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	b := strconv.Itoa(a)
	a, _ = strconv.Atoi(b)
	fmt.Println(a)
}
