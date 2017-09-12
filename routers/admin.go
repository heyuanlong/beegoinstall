package routers

import (
	"beegoinstall/controllers/admin"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var FilterUser = func(ctx *context.Context) {
    _, ok := ctx.Input.Session("username").(string)
    if !ok && ctx.Request.RequestURI != "/admin/login" {
        ctx.Redirect(302, "/admin/login")
    }
}

func init() {
	beego.InsertFilter("/admin/*",beego.BeforeRouter,FilterUser)

	beego.Router("/admin/login", &admin.LoginOutController{})
	beego.Router("/admin/logout", &admin.LoginOutController{},"*:LoginOut")

	beego.Router("/admin/admin", &admin.AdminController{})
	beego.Router("/admin/addarticle", &admin.ArticleController{})
	beego.Router("/admin/getallarticle", &admin.ArticleController{},"*:GetAllarticle")


	beego.Router("/admin/issue_article/?:id", &admin.ArticleController{},"*:Issue_article")
	beego.Router("/admin/draft_article/?:id", &admin.ArticleController{},"*:Draft_article")
	beego.Router("/admin/delete_article/?:id", &admin.ArticleController{},"*:Delete_article")
	beego.Router("/admin/update_article/?:id", &admin.ArticleController{},"*:Update_article")
	beego.Router("/admin/uploadimg", &admin.AdminController{},"*:Upload_img")


}
