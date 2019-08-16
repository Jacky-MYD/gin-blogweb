package models

import (
	db "blogWeb_gin/database"
	"blogWeb_gin/utils"
	"fmt"
)
type User struct {
	Id        int    `json:"id" form:"id"`
	UserName string `json:"username" form:"username"`
	PassWord  string `json:"password" form:"password"`
	CreateTime int64 `json:"createtime" form:"createtime"`
}

func (u * User) AddUser() (id int64, err error) {
	return utils.ModifyDB("INSERT INTO users(id, username, password, createtime) VALUES (?, ?, ?, ?)", u.Id, u.UserName, u.PassWord, u.CreateTime)

}

//按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	row := db.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

//根据用户名和密码，查询id
func QueryUserWithParam(username ,password string)int{
	sql:=fmt.Sprintf("where username='%s' and password='%s'",username,password)
	return QueryUserWightCon(sql)
}


