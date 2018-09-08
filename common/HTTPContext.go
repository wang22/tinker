package common

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/wang22/tinker/model"
)

type HTTPContext struct {
	Params  map[string]interface{}
	Context echo.Context
}

func (ctx *HTTPContext) Put(key string, value interface{}) {
	ctx.Params[key] = value
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

func (ctx *HTTPContext) Param(name string) string {
	return ctx.Context.Param(name)
}

func (ctx *HTTPContext) ParamInt(name string) int {
	value := ctx.Param(name)
	number, err := strconv.Atoi(value)
	if err != nil {
		log.Warnf("Can not convert %v to int", value)
		return -1
	}
	return number
}
