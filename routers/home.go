package routers

import (
	"beegoinstall/controllers/home"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &home.MainController{})
	beego.Router("/:id([0-9]+)", &home.MainController{})
    beego.Router("/article/?:id", &home.MainController{},"*:OneArticle")
    beego.Router("/category/:c/?:page", &home.MainController{},"*:Category")
    beego.Router("/comment", &home.MainController{},"*:Comment")
    beego.Router("/code", &home.CodeController{},"Get:GetCode")


	beego.Router("/search", &home.MainController{},"post:Search")
}
