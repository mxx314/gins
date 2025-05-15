package main

import (
	"encoding/gob"
	"fmt"
	"gins/controllers"
	"gins/filter"
	"gins/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Gin server start...")

	//初始化Gin服务
	ginServer := gin.Default()

	//初始化session
	store := cookie.NewStore([]byte("secret"))
	ginServer.Use(sessions.Sessions("session", store))
	gob.Register(models.User{})

	//注册全局中间件
	filter.RegFilter(ginServer)

	//注册路由
	controllers.Index(ginServer)
	controllers.User(ginServer)
	controllers.Json(ginServer)
	controllers.Param(ginServer)
	controllers.Redirect(ginServer)
	controllers.Course(ginServer)

	//404错误页面处理，注意：必须要将404错误页面处理放在路由注册最后
	controllers.Error(ginServer)

	//加载模版及静态资源，注意如果templates下有目录，需要定义如下格式
	ginServer.LoadHTMLGlob("templates/**/*")
	ginServer.Static("/static", "./static")

	ginServer.Run(":8089")
}
