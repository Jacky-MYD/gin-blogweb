package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//获取session
func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	loginuser := session.Get("loginuser")
	//fmt.Println("loginuser:", loginuser)
	if loginuser != nil {
		return true
	} else {
		return false
	}
}

//储存captchaId的session状态
func SetCaptchaId(c *gin.Context, captchaId string)  {
	session := sessions.Default(c)
	session.Set("CaptchaId", captchaId)
	session.Save()
	fmt.Println("123===", session.Get("CaptchaId").(string))
}

// 获取captchaId的session状态
func GetCaptchaId(c *gin.Context) string {
	session := sessions.Default(c)
	captchaId := session.Get("CaptchaId").(string)
	return captchaId
}