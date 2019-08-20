package controllers

import (
	"blogWeb_gin/models"
	"blogWeb_gin/utils"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func  LoginGet(c *gin.Context) {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)
	captchaImageUrl := "/getCaptchaPng/" + captchaId + ".png"
	SetCaptchaId(c, captchaId)

	//返回html
	c.HTML(http.StatusOK, "login.html", gin.H{"Title": "登录页", "CaptchaImageUrl": captchaImageUrl})
}

//登录
func LoginPost(c *gin.Context) {
	//获取表单信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	captchaId := GetCaptchaId(c)
	value := c.PostForm("value")

	fmt.Println("username:", username, ",password:", password, captchaId, value)

	id := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("id:", id)
	if captcha.VerifyString(captchaId, value) {
		if id > 0 {
			/*
			设置了session后悔将数据处理设置到cookie，然后再浏览器进行网络请求的时候回自动带上cookie
			因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置
			 */
			session := sessions.Default(c)
			session.Set("loginuser", username)
			session.Save()

			c.JSON(http.StatusOK, gin.H{"code": 1, "message": "登录成功"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "验证码有误"})
	}

}