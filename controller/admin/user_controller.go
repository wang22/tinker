package controller_admin

import (
	"github.com/wang22/tinker/common"
	"github.com/wang22/tinker/global"
)

type Test1 struct {
	ID   int
	Name string
}

type UserController struct {
}

func (UserController) Get(ctx common.HTTPContext) error {
	ctx.Put("sss", "sdfsdf")
	global.DB().Create(&Test1{Name: "123"})
	return ctx.JSONOK()
}
