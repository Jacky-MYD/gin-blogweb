package controllers

import (
	"blogWeb_gin/utils"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"` //验证码Id
	ImageUrl  string `json:"imageUrl"`  //验证码图片url
}

// 生成验证码
func GetCaptcha(c *gin.Context) {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)
	var captcha CaptchaResponse
	captcha.CaptchaId = captchaId
	captcha.ImageUrl = "/getCaptchaPng/" + captchaId + ".png"

	SetCaptchaId(c, captchaId)

	c.JSON(http.StatusOK, captcha)
}

// 校验验证码
func VerifyCaptcha(c *gin.Context) {
	captchaId := c.Param("captchaId")
	value := c.Param("value")

	if captchaId == "" || value == "" {
		c.String(http.StatusBadRequest, "参数错误")
	}
	if captcha.VerifyString(captchaId, value) {
		c.JSON(http.StatusOK, "验证成功")
	} else {
		c.JSON(http.StatusOK, "验证失败")
	}
}

// 生成验证码图片
func GetCaptchaPng(c *gin.Context) {
	captchaId := c.Param("captchaId")
	fmt.Println("GetCaptchaPng : " + captchaId)
	utils.ServeHTTP(c.Writer, c.Request)
}