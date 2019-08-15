package utils

import (
	"crypto/md5"
	"fmt"
	db "gin/blogWeb_gin/database"
	"log"
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


//func SwitchMarkdownToHtml(content string) template.HTML {
//	markdown := blackfriday.MarkdownCommon([]byte(content))
//
//	//获取到html文档
//	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
//	/**
//	对document进程查询，选择器和css的语法一样
//	第一个参数：i是查询到的第几个元素
//	第二个参数：selection就是查询到的元素
//	 */
//	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
//		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
//		selection.SetHtml(string(light))
//		fmt.Println(selection.Html())
//		fmt.Println("light:", string(light))
//		fmt.Println("\n\n\n")
//	})
//	htmlString, _ := doc.Html()
//	return template.HTML(htmlString)
//}

