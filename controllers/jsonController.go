package controllers

import (
	"gins/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回数据常用类型
func Json(ginServer *gin.Engine) {

	//普通数据类型
	ginServer.GET("/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})

	//map类型
	ginServer.GET("/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"name": "xiaoming", "age": 18})
	})

	//自定义map
	ginServer.GET("/json3", func(c *gin.Context) {
		m := make(map[string]interface{})
		m["name"] = "zhang3"
		m["age"] = 20
		m["sex"] = "male"
		c.JSON(http.StatusOK, m)
	})

	//结构体
	ginServer.GET("/json4", func(c *gin.Context) {
		user := models.User{
			Name: "li6",
			Age:  12,
			City: "shanghai",
		}
		c.JSON(http.StatusOK, user)
	})

	//切片结构体
	ginServer.GET("/json5", func(c *gin.Context) {
		user1 := models.User{Name: "li6", Age: 12, City: "shanghai"}
		user2 := models.User{Name: "limo", Age: 14, City: "anhui"}
		user := []models.User{user1, user2}
		c.JSON(http.StatusOK, user)
	})
}
