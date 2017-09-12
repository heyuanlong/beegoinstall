package models

import (
	"github.com/astaxie/beego/orm"
	utils "beegoinstall/utils"
	"fmt"
)

type Comment struct {
	Id       	int
	Articleid 	int
	Content 	string
	Name		string
	Time 		string
}

func init() {
	orm.RegisterModel(new(Comment))
}


// 添加文章
func AddCommet(articleid string, content string, name string ) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	time := utils.GetFormatTime("2006-01-02 15:04:05")

	sql := "insert into comment(articleid, content, name ,time  ) values(?,?,?,?)"
	res, err := o.Raw(sql, articleid, content, name,time).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}

func GetCommentByArticleId(articleid int)([]orm.Params, int64, error) {
	sql := "select name,content,time from comment where articleid =? order by time"
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql,articleid).Values(&maps)
	if err != nil {
		fmt.Println("execute sql error:"+sql)
		fmt.Println(err)
		return nil, 0, err
	}
	return maps,num,nil
}