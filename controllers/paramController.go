package controllers

import (
	"encoding/json"
	"gins/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Param(ginServer *gin.Engine) {
	//获取get请求参数
	ginServer.GET("/param1", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	//获取get请求参数 restful
	ginServer.GET("/param2/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	//from表单
	ginServer.POST("/param3", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	//json数据
	ginServer.POST("/param4", func(c *gin.Context) {
		data, _ := c.GetRawData()
		var m map[string]interface{}
		json.Unmarshal(data, &m)
		c.JSON(http.StatusOK, m)
	})

	// 解析post请求到结构体
	ginServer.POST("/param5", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, user)
		}
	})

	// 解析get请求到结构体
	ginServer.GET("/param6", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, user)
		}
	})
}
