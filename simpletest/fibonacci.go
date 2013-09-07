package main

//使用闭包输出斐波纳契数组
import (
	"fmt"
)

func fibonacci() func() int {
	var a int = 1
	var b int = 0
	var i int = 0

	return func() int {
		i++
		if i == 1 {
			return 0
		} else {
			c := a + b
			a = b
			b = c
			return c
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
