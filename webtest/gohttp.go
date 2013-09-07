package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//设置路由
	http.HandleFunc("/", sayHello)
	//启动http服务
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

//Handler Func(参数规定)
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world is version1.")
}
