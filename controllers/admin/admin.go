package admin


import (
	"github.com/astaxie/beego"
	. "beegoinstall/models"
	"crypto/md5"
    "encoding/hex"
    "fmt"
    "path"
    "time"
    "strconv"
)

type LoginOutController struct {
	beego.Controller
}

func (c *LoginOutController) Get() {
	c.TplName = "admin/login.tpl"
}

func (this *LoginOutController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	if username == "" || password == "" {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "invalid request", "refer": "/"}
	}

	user, err := FindUser(username)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "user does not exist", "refer": "/"}
	} else {
		passwd := Md5(Md5(password))

		if passwd == user.Password {
			this.SetSession("username", username)
			this.Data["json"] = map[string]interface{}{"result": true, "msg": "user[" + user.Username + "] login success ", "refer": "/admin/admin"}
		} else {
			this.Data["json"] = map[string]interface{}{"result": false, "msg": "login failed ", "refer": "/"}
		}
	}
	this.ServeJSON()
}
func (this *LoginOutController) LoginOut() {
	this.DelSession("username")
	this.Redirect("/", 302)
}







type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	username := this.GetSession("username")
	if username == nil {
		this.Redirect("/admin/login", 302)
	} else {
		this.TplName = "admin/admin.tpl"
	}

}

func (this *AdminController) Upload_img() {
    f, h, _ := this.GetFile("editormd-image-file")//获取上传的文件
    ext := path.Ext(h.Filename)
    fileSaveName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)  
    path := "static/upload/" + fileSaveName//文件目录
    f.Close()                                          //关闭上传的文件，不然的话会出现临时文件不能清除的情况

    this.SaveToFile("editormd-image-file", path)                    //存文件   
/*  	fmt.Print("----------"+path)
    fmt.Print("----------"+this.Ctx.Request.RequestURI)
    fmt.Print("----------"+this.Ctx.Request.RemoteAddr)
    fmt.Print("----------"+this.Ctx.Request.Host)*/

    this.Data["json"] = map[string]interface{}{"success": 1, "msg": "ok", "url":"/"+path}
	this.ServeJSON()
}



type ArticleController struct {
	beego.Controller
}
func (this *ArticleController) Get() {
	maps,_,_ := ListClassify()
	this.Data["classify"]=maps
	this.TplName = "admin/addarticle.tpl"
}
func (this *ArticleController) Post() {
	title 			:= this.GetString("title")
	keywords 		:= this.GetString("keywords")
	content 		:= this.GetString("content")
	status 			:= this.GetString("status")
	classifyid 		:= this.GetString("classifyid")
	articleId 		:= this.GetString("id")
	author 			:= "user"

	if articleId == "" {
		id, err := AddArticle(title, keywords, content, classifyid,author,status)
		if nil == err {
			this.Data["json"] = map[string]interface{}{
				"result": true,
				"msg":    "success added, id " + fmt.Sprintf("[%d] ", id),
				"data":   id,
				"refer":  nil,
			}
		} else {
			this.Data["json"] = map[string]interface{}{"result": false, "msg": "added failed", "refer": nil}
		}

	}else{
		ids, err := strconv.Atoi(articleId)
		if nil != err || ids < 0 {
			this.Data["json"] = map[string]interface{}{"result": false, "msg": "update failed articleId", "refer": nil}
		}else{
			err := UpdateArticle(ids,title,keywords,content,classifyid)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"result": false, "msg": "update failed UpdateArticle", "refer": nil}
			}else{
				this.Data["json"] = map[string]interface{}{
				"result": true,
				"msg":    "success updated, id " + fmt.Sprintf("[%d] ", ids),
				"data":   ids,
				"refer":  nil,
			}
			}
		}
	}
	this.ServeJSON()
}

func (this *ArticleController) GetAllarticle() {

	maps, _, err :=AllArticleAdmin()

	if nil != err {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "get list failed", "refer": "/"}
		this.ServeJSON()
	} else {
		this.Data["list"] = maps
		this.TplName = "admin/listarticle.tpl"
	}
}

func (this *ArticleController) Issue_article() {
	id := this.Ctx.Input.Param(":id")
	ids, err := strconv.Atoi(id)
	if nil != err || ids < 0 {
		
	}else{
		UpdateArticleStatus(ids,1)
	}

	this.Redirect("/admin/getallarticle",302)
}
func (this *ArticleController) Draft_article() {
	id := this.Ctx.Input.Param(":id")
	ids, err := strconv.Atoi(id)
	if nil != err || ids < 0 {
		
	}else{
		UpdateArticleStatus(ids,0)
	}

	this.Redirect("/admin/getallarticle",302)
}

func (this *ArticleController) Delete_article() {
	id := this.Ctx.Input.Param(":id")
	ids, err := strconv.Atoi(id)
	if nil != err || ids < 0 {
		
	}else{
		UpdateArticleStatus(ids,-1)
	}

	this.Redirect("/admin/getallarticle",302)
}

func (this *ArticleController) Update_article() {
	id := this.Ctx.Input.Param(":id")
	articleId, err := strconv.Atoi(id)
	if nil != err || articleId < 0 {
		this.Redirect("/admin/getallarticle",302)
		return 
	}
	art, err_ := GetArticle(articleId)
	if err_ != nil {
		this.Redirect("/admin/getallarticle",302)
		return 
	}
	this.Data["title"] = art.Title
	this.Data["keywords"] = art.Keywords
	this.Data["id"] = articleId
	this.Data["content"] = art.Content
	maps,_,_ := ListClassify()
	this.Data["classify"]=maps
	this.Data["classifyid"] = art.Classifyid
	this.TplName = "admin/addarticle.tpl"
}

func Md5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}