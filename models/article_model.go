package models

import (
	"fmt"
	"gin/blogWeb_gin/config"
	db "gin/blogWeb_gin/database"
	"gin/blogWeb_gin/utils"
	"log"
	"strconv"
)

type Article struct {
	Id         int    `json:"id" form:"id"`
	Title      string `json:"title" form:"title"`
	Tags       string `json:"tags" form:"tags"`
	Short      string `json:"short" form:"short"`
	Content    string `json:"content" form:"content"`
	Author     string `json:"author" form:"author"`
	CreateTime int64  `json:"createtime" form:"createtime"`
	//Status int //Status=0为正常，1为删除，2为冻结
}


// 插入一篇文章
func (article * Article) InsertArticle() (int64, error) {
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.CreateTime)
}

// 根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	page--
	fmt.Println("---------->page", page)
	//从配置文件中获取每页的文章数量
	return QueryArticleWithPage(page, config.NUM)
}

/**
分页查询数据库
limit分页查询语句，
    语法：limit m，n

    m代表从多少位开始获取，与id值无关
    n代表获取多少条数据

注意limit前面咩有where
 */
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	rows, err := db.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}


//----------查询文章-------------

func QueryArticleWithId(id int) Article {
	row := db.QueryRowDB("select id,title,tags,short,content,author,createtime from article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}


//------翻页------

//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var artcileRowsNum = 0

//只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	if artcileRowsNum == 0 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}

//查询文章的总条数
func QueryArticleRowNum() int {
	row := db.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

//设置页数
func SetArticleRowsNum(){
	artcileRowsNum = QueryArticleRowNum()
}

//----------修改数据----------

func UpdateArticle(article Article) (int64, error) {
	//数据库操作
	return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}

//----------删除文章---------
func DeleteArticle(artID int) (int64, error) {
	i, err := deleteArticleWithArtId(artID)
	SetArticleRowsNum()
	return i, err
}

func deleteArticleWithArtId(artID int) (int64, error) {
	return utils.ModifyDB("delete from article where id=?", artID)
}


//查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	rows, err := db.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}

//--------------按照标签查询--------------
func QueryArticlesWithTag(tag string) ([]Article, error) {

	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)
	return QueryArticlesWithCon(sql)
}