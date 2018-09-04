package controller_admin

import (
	"github.com/labstack/echo"
	"github.com/wang22/tinker/common"
)

const (
	METHOD_GET = iota
	METHOD_POST
	METHOD_PUT
	METHOD_DELETE
)

type ControllerFunc func(ctx common.HTTPContext) error

type AdminRouter struct {
	group *echo.Group
}

func (router *AdminRouter) Routing(g *echo.Group) {
	router.group = g;
	userController := new(UserController)
	router.httpGet("/get", userController.Get)
}

func (router *AdminRouter) httpMethod(method int, path string, ctrFunc ControllerFunc) {
	ctx := common.HTTPContext{}
	ctx.Param = make(map[string]interface{})
	if method == METHOD_GET {
		router.group.GET(path, func(context echo.Context) error {
			ctx.Context = context
			return ctrFunc(ctx)
		})
	} else if method == METHOD_POST {
		router.group.POST(path, func(context echo.Context) error {
			ctx.Context = context
			return ctrFunc(ctx)
		})
	} else if method == METHOD_PUT {
		router.group.PUT(path, func(context echo.Context) error {
			ctx.Context = context
			return ctrFunc(ctx)
		})
	} else if method == METHOD_DELETE {
		router.group.DELETE(path, func(context echo.Context) error {
			ctx.Context = context
			return ctrFunc(ctx)
		})
	}
}

func (router *AdminRouter) httpGet(path string, ctrFunc ControllerFunc) {
	router.httpMethod(METHOD_GET, path, ctrFunc)
}

func (router *AdminRouter) httpPost(path string, ctrFunc ControllerFunc) {
	router.httpMethod(METHOD_POST, path, ctrFunc)
}

func (router *AdminRouter) httpPut(path string, ctrFunc ControllerFunc) {
	router.httpMethod(METHOD_PUT, path, ctrFunc)
}

func (router *AdminRouter) httpDelete(path string, ctrFunc ControllerFunc) {
	router.httpMethod(METHOD_DELETE, path, ctrFunc)
}