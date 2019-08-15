package main

import (
	db "gin/blogWeb_gin/database"
	router "gin/blogWeb_gin/routers"
)

func main()  {
	db.InitMysql()

	// 路由
	router := router.InitRouter()

	// 静态资源
	router.Static("/static", "./blogWeb_gin/static")


	router.Run(":8081")
}
