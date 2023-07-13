package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-service/pkg/assert"
	"go-service/pkg/errno"
	"go-service/pkg/manager"
	"go-service/pkg/restapi"
	"go-service/work/app/biz/application/app"
	"go-service/work/app/biz/application/cqe"
	"sync"
)

var (
	userControllerOnce      sync.Once
	singletonUserController UserController
)

type UserController interface {
	manager.Controller
	CreateUser(c *gin.Context)
}

// UserControllerPlugin 用于初始化UserController
type UserControllerPlugin struct {
}

func (p *UserControllerPlugin) Name() string {
	return "UserControllerPlugin"
}

func (p *UserControllerPlugin) MustCreateController() manager.Controller {
	return DefaultUserController()
}

// DefaultUserController 返回UserController单例
func DefaultUserController() UserController {
	assert.NotCircular()
	userControllerOnce.Do(func() {
		singletonUserController = &userControllerImpl{
			userApp: app.DefaultUserApp(),
		}
	})
	assert.NotNil(singletonUserController)
	return singletonUserController
}

type userControllerImpl struct {
	userApp app.UserApp
}

func (ctrl *userControllerImpl) RegisterOpenApi(group *gin.RouterGroup) {
	g := group.Group("/user")
	{
		g.POST("/create", ctrl.CreateUser)
	}
}

func (ctrl *userControllerImpl) RegisterInnerApi(group *gin.RouterGroup) {
}

func (ctrl *userControllerImpl) RegisterDebugApi(group *gin.RouterGroup) {
}

func (ctrl *userControllerImpl) CreateUser(c *gin.Context) {
	cmd := cqe.CreateAccountCommand{}
	if err := c.BindJSON(&cmd); err != nil {
		restapi.Failed(c, errno.NewSimpleBizError(errno.ErrParameterInvalid, err, "body"))
		return
	}

	m, err := ctrl.userApp.CreateUser(context.Background(), &cmd)
	if err != nil {
		restapi.Failed(c, err)
		return
	}

	restapi.Success(c, m)
}
