package controllers

import (
	"gins/filter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Course(ginServer *gin.Engine) {
	//路由器定义，下面的路由需要使用courseGroup
	courseGroup := ginServer.Group("/course")
	{
		// get请求
		courseGroup.GET("/list", filter.MyHandler(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "course list"})
		})

		// post请求
		courseGroup.POST("/save", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "course save"})
		})

		// update请求
		courseGroup.PUT("/update", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "course update"})
		})

		// delete请求
		courseGroup.DELETE("/delete", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "course delete"})
		})
	}

}
