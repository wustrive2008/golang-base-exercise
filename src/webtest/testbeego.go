package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Ctx.WriteString("hello wrold!")
}
func main() {

	beego.Router("/", &HomeController{})
	beego.Run()
}
