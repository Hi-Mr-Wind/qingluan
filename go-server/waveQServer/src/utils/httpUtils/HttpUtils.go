package httpUtils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"waveQServer/src/comm"
	"waveQServer/src/core/cache"
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
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token,session,XMD-TOKEN,API_KEY")
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
	parsedToken, err := jwtutil.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusForbidden, comm.Fail(err.Error()))
		c.Abort()
		return
	}
	expirationTime := parsedToken.ExpiresAt
	currentTime := time.Now().Unix()
	if expirationTime-currentTime <= 300 {
		renewedToken, renewErr := jwtutil.RenewToken(parsedToken) // 生成新token
		if renewErr != nil {
			c.JSON(http.StatusForbidden, comm.Fail(renewErr.Error()))
			c.Abort()
			return
		}
		if len(renewedToken) == 0 { // 如果没有超时则结束本方法
			return
		}
		c.Header("XMD-TOKEN", renewedToken) // 更新新token
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
	//如果apikey并不存在
	if !cache.IsInKeys(apiKey) {
		c.JSON(http.StatusForbidden, comm.Fail("Unknown client!"))
		c.Abort()
		return
	}
	user := cache.GetUser(apiKey)
	//用户期限是否是无限期
	if user.ExpirationTime != -1 {
		//如果api已经过期则删除
		if user.ExpirationTime <= time.Now().UnixMilli() {
			cache.DelApiKey(user.ApiKey)
			c.JSON(http.StatusForbidden, comm.Fail("This key has expired"))
			c.Abort()
			return
		}
	}
}
