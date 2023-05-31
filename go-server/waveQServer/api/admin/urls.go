package admin

import "github.com/gin-gonic/gin"

func Enter(url *gin.RouterGroup) {
	url.POST("/login", Login)
}

func Urls(url *gin.RouterGroup) {
	url.GET("/login", Login)
}
