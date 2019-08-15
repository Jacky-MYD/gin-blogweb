package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gin/blogWeb_gin/models"
	"net/http"
	"strconv"
)

// 显示文章
func ShowArticleGet(c *gin.Context) {

	//获取session
	islogin := GetSession(c)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:", id)

	//获取id所对应的文章信息
	art := models.QueryArticleWithId(id)
	//渲染HTML
	c.HTML(http.StatusOK, "show_article.html", gin.H{"IsLogin": islogin,"Title":art.Title,"Content":art.Content})
}
