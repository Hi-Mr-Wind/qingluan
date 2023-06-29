package admin

import (
	"github.com/gin-gonic/gin"
	"waveQServer/src/service"
)

// Enter 管理员登录
func Enter(url *gin.RouterGroup) {
	url.POST("/login", service.Login)
}

func Urls(url *gin.RouterGroup) {
	url.POST("/createGroup", service.CreateGroup)   //创建组对象
	url.POST("/createApiKey", service.CreateApiKey) //创建消费者apikey
	url.POST("/create_token", service.CreateToken)  // 创建令牌
	url.POST("/delete_token", service.DeleteToken)  // 删除令牌
}
