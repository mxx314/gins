package filter

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// 中间件放行与挂起
func MyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("core start....")
		if true {
			c.Next()
		} else {
			c.Abort()
		}

		fmt.Println("core end...")
	}
}

// 中间件放行与挂起
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := "xiaxia"
		c.Set("username", username)
		if username == "xiaxia" {
			fmt.Println("身份验证成功！！！")
			c.Next()
		} else {
			c.Abort()
			c.JSON(200, gin.H{"message": "身份验证失败！"})
		}
	}
}

// 中间件全局注册
func RegFilter(ginServer *gin.Engine) {
	//ginServer.Use(myHandler())
	ginServer.Use(LoginHandler())
	ginServer.Use(Cors())
}

// 跨域请求配置
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	},
	)
}
