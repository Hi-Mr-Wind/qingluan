package http

import (
	"sync"
)

var (
	userControllerOnce      sync.Once
	singletonUserController UserController
)

type UserController interface {
	manager.Controller
}
