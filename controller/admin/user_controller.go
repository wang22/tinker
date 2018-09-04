package controller_admin

import (
	"fmt"
	"github.com/wang22/tinker/common"
)

type UserController struct {
}

func (UserController) Get(ctx common.HTTPContext) error {
	fmt.Println("123")
	ctx.Put("sss", "sdfsdf")
	return ctx.String(200, "ddd")
}

