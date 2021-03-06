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
	return ctx.Context.JSON(http.StatusOK, model.HTTPJSONResult{http.StatusOK, "ok", ctx.Params})
}

func (ctx *HTTPContext) JSONErr(msg string) error {
	if msg == "" {
		msg = "err"
	}
	return ctx.Context.JSON(http.StatusInternalServerError, model.HTTPJSONResult{http.StatusInternalServerError, msg, ctx.Params})
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

func (ctx *HTTPContext) Bind(i interface{}) error {
	return ctx.Context.Bind(i)
}

func (ctx *HTTPContext) SetCookie(cookie *http.Cookie) {
	ctx.Context.SetCookie(cookie)
}
