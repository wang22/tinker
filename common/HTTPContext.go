package common

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/wang22/tinker/model"
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
	return ctx.Context.JSON(http.StatusOK, model.HTTPJSONResult{http.StatusOK, "ok", ctx.Param})
}

func (ctx *HTTPContext) JSONErr() error {
	return ctx.Context.JSON(http.StatusOK, model.HTTPJSONResult{http.StatusInternalServerError, "error", ctx.Param})
}
