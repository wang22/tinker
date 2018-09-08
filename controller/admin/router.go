package controller_admin

import (
	"github.com/labstack/echo"
	"github.com/wang22/tinker/common"
)

type ControllerFunc func(ctx common.HTTPContext) error

type AdminRouter struct {
	group *echo.Group
}

func (router *AdminRouter) Routing(g *echo.Group) {
	router.group = g
	adminController := new(AdminController)
	router.httpPost("/auth/login", adminController.Login)
}

func (router *AdminRouter) makeContext(context echo.Context) common.HTTPContext {
	ctx := common.HTTPContext{}
	ctx.Params = make(map[string]interface{})
	ctx.Context = context
	return ctx
}

func (router *AdminRouter) httpGet(path string, ctrFunc ControllerFunc) {
	router.group.GET(path, func(context echo.Context) error {
		return ctrFunc(router.makeContext(context))
	})
}

func (router *AdminRouter) httpPost(path string, ctrFunc ControllerFunc) {
	router.group.POST(path, func(context echo.Context) error {
		return ctrFunc(router.makeContext(context))
	})
}

func (router *AdminRouter) httpPut(path string, ctrFunc ControllerFunc) {
	router.group.PUT(path, func(context echo.Context) error {
		return ctrFunc(router.makeContext(context))
	})
}

func (router *AdminRouter) httpDelete(path string, ctrFunc ControllerFunc) {
	router.group.DELETE(path, func(context echo.Context) error {
		return ctrFunc(router.makeContext(context))
	})
}
