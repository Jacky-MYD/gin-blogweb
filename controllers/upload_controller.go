package controllers

import (
	"fmt"
	"blogWeb_gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadPost(c *gin.Context) {
	fmt.Println("fileupload...")

	fileHeader, err := c.FormFile("upload")
	if err != nil {
		responseErr(c, err)
		return
	}
	fmt.Println("name:", fileHeader.Filename, fileHeader.Size)

	now := time.Now()
	fmt.Println("ext:", filepath.Ext(fileHeader.Filename))
	fileType := "other"
	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	//文件夹路径存储到项目的路径
	fileDir := fmt.Sprintf("blogWeb_gin/static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//文件夹路径存储到数据库的路径，以便于显示
	fileDb := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		responseErr(c, err)
		return
	}

	fmt.Println("fileDir:", fileDir)
	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	fileDbPathStr := filepath.Join(fileDb, fileName)


	//将浏览器客户端上传的文件拷贝到本地路径的文件里面，此处也可以使用io操作
	c.SaveUploadedFile(fileHeader,filePathStr)

	if fileType == "img" {
		album := models.Album{0, fileDbPathStr, fileName, 0, timeStamp}
		album.InsertAlbum()
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "上传成功"})

}

func responseErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
}
