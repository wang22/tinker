package controller_admin

import (
	"github.com/labstack/echo"
	"github.com/wang22/tinker/common"
)

const (
	MethodGet = iota
	MethodPost
	MethodPut
	MethodDelete
)

type ControllerFunc func(ctx common.HTTPContext) error

type AdminRouter struct {
	group *echo.Group
}

func (router *AdminRouter) Routing(g *echo.Group) {
	router.group = g
	// userController := new(UserController)
	// router.httpGet("/get", userController.Get)
	// router.httpGet("/get/:id", userController.Get1)

	adminController := new(AdminController)
	router.httpPost("/auth/login", adminController.Login)
}

func (router *AdminRouter) httpMethod(method int, path string, ctrFunc ControllerFunc) {
	if method == MethodGet {
		router.group.GET(path, func(context echo.Context) error {
			ctx := common.HTTPContext{}
			ctx.Params = make(map[string]interface{})
			ctx.Context = context
			return ctrFunc(ctx)
		})
	} else if method == MethodPost {
		router.group.POST(path, func(context echo.Context) error {
			ctx := common.HTTPContext{}
			ctx.Params = make(map[string]interface{})
			ctx.Context = context
			return ctrFunc(ctx)
		})
	} else if method == MethodPut {
		router.group.PUT(path, func(context echo.Context) error {
			ctx := common.HTTPContext{}
			ctx.Params = make(map[string]interface{})
			ctx.Context = context
			return ctrFunc(ctx)
		})
	} else if method == MethodDelete {
		router.group.DELETE(path, func(context echo.Context) error {
			ctx := common.HTTPContext{}
			ctx.Params = make(map[string]interface{})
			ctx.Context = context
			return ctrFunc(ctx)
		})
	}
}

func (router *AdminRouter) httpGet(path string, ctrFunc ControllerFunc) {
	router.httpMethod(MethodGet, path, ctrFunc)
}

func (router *AdminRouter) httpPost(path string, ctrFunc ControllerFunc) {
	router.httpMethod(MethodPost, path, ctrFunc)
}

func (router *AdminRouter) httpPut(path string, ctrFunc ControllerFunc) {
	router.httpMethod(MethodPut, path, ctrFunc)
}

func (router *AdminRouter) httpDelete(path string, ctrFunc ControllerFunc) {
	router.httpMethod(MethodDelete, path, ctrFunc)
}
