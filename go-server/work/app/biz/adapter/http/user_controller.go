package http

import (
	"github.com/Hi-Mr-Wind/qingluan/tree/develop-ddd/go-server/pkg/manager"
	"go-service/pkg/assert"
	"sync"
)

var (
	userControllerOnce      sync.Once
	singletonUserController UserController
)

type UserController interface {
	manager.Controller
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
		singletonUserController = &userControllerImpl{}
	})
	assert.NotNil(singletonUserController)
	return singletonUserController
}

type userControllerImpl struct {
}
