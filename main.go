package main

import (
	"blogWeb_gin/utils"
	db "blogWeb_gin/database"
	router "blogWeb_gin/routers"
)

func main()  {

	db.InitMysql()
	// 路由
	router := router.InitRouter()

	// 注册一个全局中间件
	router.Use(utils.StatCost())

	// 静态资源
	router.Static("/static", "./static")


	router.Run(":8081")
}
