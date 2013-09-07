/**
* 数组切片：大小可自动扩展
*
* 创建数组切片：
*
 */
package main

import (
	"fmt"
)

func main() {

	/* 基于数组创建

	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[5:]

	fmt.Println("Elements of myArray:")

	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice")

	for _, v := range mySlice {

		fmt.Print(v, " ")
	}

	fmt.Println()
	*/
	//创建一个初始元素个数为5的切片，元素初始值为0
	//mySlice1 := make([]int, 5)
	//创建一个初始元素个数为5的切片，元素初始值为0,预留10元素存储空间
	mySlice2 := make([]int, 5, 10)
	//直接创建
	mySlice3 := []int{1, 2, 3, 4, 5}

	mySlice2 = append(mySlice2, 1, 2, 3)

	mySlice2 = append(mySlice2, mySlice3...)

	mySlice4 := mySlice3[4:]

	fmt.Println("len(mySlice2):", len(mySlice2))

	fmt.Println("cap(mySlice2):", cap(mySlice2))

	copy(mySlice4, mySlice3)

	for _, v := range mySlice4 {
		fmt.Println(v, " ")
	}

	for _, v := range mySlice3 {
		fmt.Print(v, " ")
	}

}
