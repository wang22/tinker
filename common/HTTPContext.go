package common

import "github.com/labstack/echo"

type HTTPContext struct {
	Param map[string]interface{}
	Context echo.Context
}

func (ctx *HTTPContext) Put(key string, value interface{}) {
	ctx.Param[key] = value
}

func (ctx *HTTPContext) String(code int, s string) error {
	return ctx.String(code, s)
}

func (ctx *HTTPContext) JSON(code int, i interface{}) error {
	return ctx.JSON(code, i)
}