package home

import (
	"github.com/astaxie/beego"
	. "beegoinstall/models"
	"strconv"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
		this.TplName = "index.tpl"
}

func (this *MainController) Search() {
	data := this.GetString("data")

	this.Data["value"] = data
	this.TplName = "search.tpl"
	this.Data["list"] = make([]interface{},0)

	if len(data) == 0{
		return
	}
	maps, _, err :=ListTitleLike(data)
	if nil != err {
		return
	}

	this.Data["list"] = maps
	this.TplName = "search.tpl"
}

func (this *MainController) OneArticle() {
	var articleId int
	id := this.Ctx.Input.Param(":id")
	id_, err := strconv.Atoi(id)
	if nil != err || id_ < 0 {
		articleId = 1
	} else {
		articleId = id_
	}

	art, err_ := GetArticle(articleId)
	if err_ != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "get OneArticle failed", "refer": "/"}
		this.ServeJSON()
		return 
	}
	UpdateArticleCount(articleId)
	this.Data["id"] = art.Id
	this.Data["title"] = art.Title
	this.Data["time"] = art.Time
	this.Data["author"] = art.Author
	this.Data["count"] = art.Count
	this.Data["content"] = art.Content
	this.Data["classifyid"] = art.Classifyid

	hottest,err:= HottestArticleList(15)
	if nil == err{
		this.Data["hottest"] = hottest
	}


	classify,err_:= ClassArticleList()
	if nil == err_{
		this.Data["classify"] = classify
	}
	for _,v2 := range classify{
		cid, _ := strconv.Atoi(  v2["classifyID"].(string))
   		if art.Classifyid == cid {
   			this.Data["classifyname"] = v2["name"]
   			break
   		}
   	}

   	maps, _, _ := GetCommentByArticleId(articleId)
   	this.Data["commentlist"] = maps

	this.TplName = "home/onearticle.tpl"

}

func (this *MainController) Category() {
	var page int
	var classify int
	var prev_page_flag bool
	var next_page_flag bool
	perPageNums := 20

	c := this.Ctx.Input.Param(":c")
	cTmp, err := strconv.Atoi(c)
	if nil != err || cTmp < 0 {
		this.Redirect("/",302)
		return 
	} else {
		classify = cTmp
	}

	pageTmp := this.Ctx.Input.Param(":page")
	pageParm, err := strconv.Atoi(pageTmp)
	if nil != err || pageParm < 0 {
		page = 1
	} else {
		page = pageParm
	}

	allnums ,err_ := AllArticleCout(classify)
	if err_ != nil {
	 	allnums = 0
	 } 

	maps, _, err__ := ListArticle(page,perPageNums,classify)
	if nil != err__ {
		this.Redirect("/",302)
		return
	} else {
		if page > 1 {
			prev_page_flag = true
			this.Data["prev_page_flag"] = prev_page_flag
			this.Data["prev_page"] = fmt.Sprintf("/category/%d/%d",classify, page-1)
		}
		if (page * perPageNums) < allnums {
			next_page_flag = true
			this.Data["next_page_flag"] = next_page_flag
			this.Data["next_page"] = fmt.Sprintf("/category/%d/%d",classify, page+1)
		}

		hottest,err:= HottestArticleList(15)
		if nil == err{
			this.Data["hottest"] = hottest
		}
		classify,err_:= ClassArticleList()
		if nil == err_{
			this.Data["classify"] = classify
		}
		for _, v := range maps{
		   	for _,v2 := range classify{
		   		if v["classifyID"] == v2["classifyID"] {
		   			v["classifyname"] =  v2["name"]
		   		}
		   	}
		}

		this.Data["list"] = maps
		this.TplName = "home/categorylist.tpl"
	}

}


func (this *MainController) Comment() {
	articleid := this.GetString("articleid")
	name := this.GetString("name")
	content := this.GetString("content")
	code := this.GetString("code")

	vercode := this.GetSession("code")
	this.DelSession("code")

	if vercode != code {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "code error", "refer": "/article/"+articleid }
		this.ServeJSON()
		return 
	}
	_, err := AddCommet(articleid,content,name)
	if nil == err {
		this.Data["json"] = map[string]interface{}{
			"result": true,
			"msg":    "success comment",
			"refer":  "/article/"+articleid,
		}
	} else {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "added failed", "refer": "/article/"+articleid}
	}
	this.ServeJSON()
}