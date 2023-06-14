package admin

import "github.com/gin-gonic/gin"

// Enter 管理员登录
func Enter(url *gin.RouterGroup) {
	url.POST("/login", Login)
}

func Urls(url *gin.RouterGroup) {
	url.POST("/createGroup", CreateGroup) //创建组对象

}
