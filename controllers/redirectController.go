package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向
func Redirect(ginServer *gin.Engine) {

	// 站外重定向
	ginServer.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 站内重定向
	ginServer.GET("/test1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})
}
