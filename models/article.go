package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
	utils "beegoinstall/utils"
)

type Article struct {
	Id       		int
	Title    		string
	Keywords 		string
	Content  		string
	Author   		string
	Time     		string
	Count    		int
	Classifyid 		int
	Status   		int
}

func init() {
	orm.RegisterModel(new(Article))
}

func GetArticle(id int) (Article, error) {
	var err error
	var art Article

	o := orm.NewOrm()
	o.Using("default")
	art = Article{Id: id}
	err = o.Read(&art, "id")

	return art, err
}
func UpdateArticleCount(id int) () {
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("update article set count = count +1 where id=?",id).Exec()
}
func UpdateArticleStatus(id int,status int) () {
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("update article set status = ? where id=?",status,id).Exec()
}
func UpdateArticle(id int,title string,  keywords string,content string,classifyid string ) (error) {
	o := orm.NewOrm()
	o.Using("default")
	_ ,err := o.Raw("update article set title = ?,keywords=?,content=?,classifyID=? where id=?",title,keywords,content,classifyid,id).Exec()
	return err
}


// 添加文章
func AddArticle(title string,  keywords string,content string, classifyid string ,author string, status string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	time := utils.GetFormatTime("2006-01-02 15:04:05")

	sql := "insert into article(title, keywords, content, author ,time,count,classifyID,status ) values(?, ?, ?, ?, ?,?,?,?)"
	res, err := o.Raw(sql, title, keywords, content, author,time,1,classifyid,status).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}

func AllArticle()([]orm.Params, int64, error) {
	sql := "select * from article where status != -1 order by time desc"
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Println("execute sql1 error:"+sql)
		fmt.Println(err)
		return nil, 0, err
	}
	return maps,num,nil
}
func AllArticleAdmin()([]orm.Params, int64, error) {
	sql := "select id,title,time,status from article where status != -1 order by time desc"
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Println("execute sql1 error:"+sql)
		fmt.Println(err)
		return nil, 0, err
	}
	return maps,num,nil
}

func AllArticleCout(c int)( int, error) {
	var sql string
	if c < 0 {
		c=1
		sql = "select count(*) as number  from article where status = 1 and 1=?"
	}else{
		sql = "select count(*) as number  from article where status = 1 and classifyID=?"
	}
	
	var maps []orm.Params
	o := orm.NewOrm()
	_, err := o.Raw(sql,c).Values(&maps)
	if err != nil {
		fmt.Println("execute sql error:"+sql)
		fmt.Println(err)
		return 0, err
	}
	return strconv.Atoi(maps[0]["number"].(string))
}

func ListArticle(page int,numPerPage int,classify int)([]orm.Params, int64, error) {
	var sql string
	if page <=0 {
		page = 1
	}
	if classify < 0 {
		classify = 1
		sql = "select * from article where status = 1 and 1=? order by time desc limit ?," + fmt.Sprintf("%d", numPerPage)
	}else{
		sql = "select * from article where status = 1 and classifyID=? order by time desc limit ?," + fmt.Sprintf("%d", numPerPage)
	}

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql,classify, numPerPage*(page-1)).Values(&maps)
	if err != nil {
		fmt.Println("execute sql error:" + sql)
		fmt.Println(err)
		return nil, 0, err
	}

	return maps,num,nil
}
func ListClassify()([]orm.Params, int64, error) {
	sql := "select id,name from classify"
	var maps []orm.Params
	o := orm.NewOrm()
	num, err_ := o.Raw(sql).Values(&maps)
	if err_ != nil {
		fmt.Println("execute sql error:" + sql)
		fmt.Println(err_)
		return nil, 0, err_
	}
	return maps,num,nil
}

// 最热文章列表 
func HottestArticleList(num int) ([]orm.Params, error) {
	var maps []orm.Params

	sql := "select id,title,count from article where status = 1 order by count desc limit ?"
	o := orm.NewOrm()
	_, err := o.Raw(sql,num).Values(&maps)

	return maps, err
}
// 文章分类列表 
func ClassArticleList() ([]orm.Params, error) {
	var maps []orm.Params

	sql := "select a.classifyID,c.name ,count(*) as count from  article as a  left join classify as c on a.classifyID=c.id group by classifyID"
	o := orm.NewOrm()
	_, err := o.Raw(sql).Values(&maps)

	return maps, err
}


//--------------------------------------------------------------------------------------------

func ListTitleLike(str string)([]orm.Params, int64, error) {
	var maps []orm.Params
	if len(str) == 0 {
		return maps,0,nil
	}
	var sql string
	sql = "select id,title,keywords,author,time,classifyid from article where status = 1 and title like '%"+str+"%' limit 20;"

	o := orm.NewOrm()
	num, err := o.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Println("execute sql error:" + sql)
		fmt.Println(err)
		return nil, 0, err
	}
	return maps,num,nil
}