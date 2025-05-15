package controllers

import (
	"gins/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ginServer *gin.Engine) {

	// 返回map消息
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "hello world!"})
	})

	// 返回html页面，并将数据加载到页面中
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{"message": "hello world!", "name": "mumu"})
	})

	// 返回html页面，并将数据加载到页面中
	ginServer.GET("/index1", func(context *gin.Context) {
		context.HTML(200, "index1.html", gin.H{"message": "hello world!1", "name": "mumu1"})
	})

	// 返回html页面，并将数据加载到页面中，使用中间件传递数据，上游中间件将数据写入其中，下游中间件可以拿到
	ginServer.GET("/index2", func(context *gin.Context) {
		context.HTML(200, "index1.html", gin.H{"message": "hello world!1", "name": context.GetString("username")})
	})

	// 登陆方法，并将数据存储到session中
	ginServer.GET("/login", func(c *gin.Context) {
		var user models.User
		session := sessions.Default(c)
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			session.Set("user", user)
			session.Set("id", 100)
			session.Save()
			c.JSON(http.StatusOK, user)
		}
	})
}
