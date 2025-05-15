package controllers

import (
	"gins/filter"
	"gins/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func User(ginServer *gin.Engine) {
	// get请求
	ginServer.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "get user"})
	})

	// post请求
	ginServer.POST("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "post user"})
	})

	// put请求
	ginServer.PUT("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "put user"})
	})

	// delete请求
	ginServer.DELETE("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "delete user"})
	})

	// session获取其他路由保存的信息
	ginServer.GET("/user/detail", func(c *gin.Context) {
		session := sessions.Default(c)
		data := session.Get("user")
		user := data.(models.User)
		//user := get.(models.User)
		c.JSON(200, user)
	})

	// 为跨域提供数据
	ginServer.GET("/user/detail1", func(c *gin.Context) {
		user := models.User{
			Name: "admin",
			Age:  20,
			City: "tianjin",
		}
		//user := get.(models.User)
		c.JSON(200, user)
	})

	// 模拟跨域请求
	ginServer.GET("/user/vue", filter.Cors(), func(c *gin.Context) {
		c.HTML(200, "vue.html", gin.H{"message": "vue"})
	})
}
