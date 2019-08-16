package controllers

import (
	"fmt"
	"blogWeb_gin/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//点击删除后重定向到首页
func  DeleteArticleGet(c *gin.Context) {

	idstr := c.Query("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println("删除 id:", id)



	_, err := models.DeleteArticle(id)
	if err != nil {
		log.Println(err)
	}
	//c.JSON(http.StatusOK, gin.H{"IsLogin": islogin})
	c.Redirect(http.StatusMovedPermanently,"/")
}