package main

import (
	_ "beegoinstall/routers"
	_ "beegoinstall/initial"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

