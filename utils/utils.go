package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	db "blogWeb_gin/database"
	"github.com/PuerkitoBio/goquery"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/sourcegraph/syntaxhighlight"
	"gopkg.in/gomail.v2"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

//传入的数据不一样，那么MD5后的32位长度的数据肯定会不一样
func MD5(str string) string{
	md5str:=fmt.Sprintf("%x",md5.Sum([]byte(str)))
	return md5str
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.SqlDB.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}


//将传入的时间戳转为时间
func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}


func SwitchMarkdownToHtml(content string) template.HTML {
	markdown := blackfriday.Run([]byte(content))

	//获取到html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	/**
	对document进程查询，选择器和css的语法一样
	第一个参数：i是查询到的第几个元素
	第二个参数：selection就是查询到的元素
	 */
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
		fmt.Println("\n\n\n")
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("UserName", "Jacky")

		// 执行其他中间件
		c.Next()
		// 计算耗时
		cost := time.Since(start)
		log.Println("cost========",cost)
	}
}

// 发送邮件
func sendEmail()  {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "*****@qq.com", "Jacky") // 发件人
	m.SetHeader("To", // 收件人
		m.FormatAddress("*****@163.com", "乔峰"),
		m.FormatAddress("*****@qq.com", "郭靖"),
	)
	m.SetHeader("Subject", "Gomail")                                                            // 主题
	m.SetBody("text/html", "Hello <a href = \"https://github.com/Jacky-MYD/gin-blogweb\">git.com</a>") // 正文

	d := gomail.NewDialer("smtp.qq.com", 465, "*****@qq.com", "*********") // 发送邮件服务器、端口、发件人账号、发件人授权码
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("发送完成！")
}

// 图片验证码
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	fmt.Println("file : " + file)
	fmt.Println("ext : " + ext)
	fmt.Println("id : " + id)
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("reload : " + r.FormValue("reload"))
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}