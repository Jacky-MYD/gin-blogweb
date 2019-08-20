package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func  AboutMeGet(c *gin.Context) {

	//获取session
	islogin := GetSession(c)

	c.HTML(http.StatusOK, "aboultme.html", gin.H{"IsLogin": islogin,"wechat":"微信：miao657414127","qq":"QQ：657414127","tel":"Tel：13666666666"})
}
