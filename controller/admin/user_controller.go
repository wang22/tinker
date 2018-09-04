package controller_admin

import (
	"github.com/wang22/tinker/common"
)

type UserController struct {
}

func (UserController) Get(ctx common.HTTPContext) error {
	ctx.Put("sss", "sdfsdf")
	return ctx.JSONOK()
}
