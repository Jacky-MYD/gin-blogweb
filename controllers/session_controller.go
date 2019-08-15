package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//获取session
func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	fmt.Println("session========:", session)
	loginuser := session.Get("loginuser")
	fmt.Println("loginuser:", loginuser)
	if loginuser != nil {
		return true
	} else {
		return false
	}
}
