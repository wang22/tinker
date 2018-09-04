package common

import (
	"github.com/labstack/echo"
	"github.com/wang22/tinker/model"
	"net/http"
)

type HTTPContext struct {
	Param   map[string]interface{}
	Context echo.Context
}

func (ctx *HTTPContext) Put(key string, value interface{}) {
	ctx.Param[key] = value
}

func (ctx *HTTPContext) String(code int, s string) error {
	return ctx.String(code, s)
}

func (ctx *HTTPContext) JSON(code int, i interface{}) error {
	return ctx.Context.JSON(code, i)
}

func (ctx *HTTPContext) JSONOK() error {
	result := model.HTTPJSONResult{http.StatusOK, "ok", ctx.Param}
	return ctx.Context.JSON(http.StatusOK, result)
}

func (ctx *HTTPContext) JSONErr() error {
	result := model.HTTPJSONResult{http.StatusInternalServerError, "error", ctx.Param}
	return ctx.Context.JSON(http.StatusOK, result)
}
