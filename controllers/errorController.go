package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(ginServer *gin.Engine) {

	// 404错误处理
	ginServer.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
}
