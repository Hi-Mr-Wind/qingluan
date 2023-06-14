package httpUtils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"waveQServer/src/comm"
	"waveQServer/src/core/database"
	"waveQServer/src/core/database/dto"
	"waveQServer/src/utils"
	"waveQServer/src/utils/jwtutil"
	"waveQServer/src/utils/logutil"
)

var server *gin.Engine

// GetServer 获取一个服务引擎
func GetServer() *gin.Engine {
	return server
}

// 初始化gin服务
func init() {
	if server == nil {
		gin.DefaultWriter = logutil.GetLogFileIo() //输出日志切换为文件
		server = gin.New()
		server.Use(gin.Recovery())
		server.Use(CrosHandler())
	}
}

// GetRouterGroup 获取一个路由组
func GetRouterGroup(groupUrl string) *gin.RouterGroup {
	return server.Group(groupUrl)
}

// CrosHandler 配置跨域中间件
func CrosHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				logutil.LogError(fmt.Sprintf("%v", err))
			}
		}()
		c.Next()
	}
}

// Token 验证Token
func Token(c *gin.Context) {
	token := c.Request.Header.Get("XMD-TOKEN")
	if utils.IsEmpty(token) {
		c.JSON(http.StatusForbidden, comm.Fail("illegal user ！"))
		c.Abort()
		return
	}
	_, err := jwtutil.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusForbidden, comm.Fail(err.Error()))
		c.Abort()
		return
	}
	return
}

// VerifyUser 检查请求中的key的合法性
func VerifyUser(c *gin.Context) {
	apiKey := c.Request.Header.Get("API_KEY")
	if utils.IsEmpty(apiKey) {
		c.JSON(http.StatusForbidden, comm.Fail("Unknown client!"))
		c.Abort()
		return
	}
	user := new(dto.User)
	database.GetDb().Find(&user, "api_key = ?", apiKey)
	parse, err := time.Parse("2006-01-02 15:04:05", user.ExpirationTime)
	if err != nil {
		logutil.LogError(err.Error())
		return
	}
	if utils.IsEmpty(user.ExpirationTime) {
		return
	}
	if parse.UnixNano() < time.Now().UnixNano() {
		c.JSON(http.StatusForbidden, comm.Fail("This key has expired"))
	}
}