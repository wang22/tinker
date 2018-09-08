package controller_admin

import (
	"fmt"
	"net/http"

	"github.com/wang22/tinker/config"
	"github.com/wang22/tinker/model"

	"github.com/labstack/echo"
	"github.com/wang22/tinker/common"
)

type ControllerFunc func(ctx common.HTTPContext) error

type AdminRouter struct {
	group *echo.Group
}

func (router *AdminRouter) baseAuth() {
	router.group.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			uri := c.Request().RequestURI
			if uri == config.Get(config.ConstHttpAdminPrefix)+"/auth/login" {
				return next(c)
			}
			errModel := model.HTTPJSONResult{Code: http.StatusUnauthorized, Msg: "请先登录"}
			cookie, err := c.Cookie(config.Get(config.ConstAuthCookieKey))
			if err != nil {
				return c.JSON(http.StatusUnauthorized, errModel)
			}
			// if cookie.Value == ?? {} 验证
			fmt.Println(cookie)
			return next(c)
		}
	})
}

func (router *AdminRouter) Routing(g *echo.Group) {
	router.group = g
	router.baseAuth()

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
