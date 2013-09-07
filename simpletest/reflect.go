package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}

func main() {
	s := "this is string"
	fmt.Println(reflect.TypeOf(s)) //string
	fmt.Println("------------------")

	fmt.Println(reflect.ValueOf(s)) //this is string
	var x float64 = 3.4
	fmt.Println(reflect.ValueOf(x)) //<float64 Value>
	fmt.Println("------------------")

	a := new(MyStruct)
	a.name = "hello world"
	typ := reflect.TypeOf(a)
	fmt.Println(typ.NumMethod()) //1
	fmt.Println("------------------")

	b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(b[0]) //hello world
	fmt.Println("-------------------")

	var aa MyStruct
	bb := new(MyStruct)

	fmt.Println(reflect.ValueOf(aa)) //<main.MyStruct Value>
	fmt.Println(reflect.ValueOf(bb)) //<*main.MyStruct Value>

	fmt.Println("--------------------")
	aa.name = "hello world"
	bb.name = "hello world"
	val := reflect.ValueOf(aa).FieldByName("name")

	fmt.Println(val) //hello world

	fmt.Println("--------------------")
	fmt.Println(reflect.ValueOf(aa).FieldByName("name").CanSet()) //false
	fmt.Println(reflect.ValueOf(&(aa.name)).Elem().CanSet())      //true

	fmt.Println("-------------------")
	var c string = "hello world"
	p := reflect.ValueOf(&c)
	fmt.Println(p.CanSet())        //false
	fmt.Println(p.Elem().CanSet()) //true
	p.Elem().SetString("newName")
	fmt.Println(c) //newName
}
