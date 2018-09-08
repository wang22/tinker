package controller_admin

import (
	"github.com/wang22/tinker/common"
	"github.com/wang22/tinker/global"
	"github.com/wang22/tinker/model"
)

type UserController struct {
}

func (UserController) Get(ctx common.HTTPContext) error {
	ctx.Put("sss", "sdfsdf")
	// global.DB().Create(&model.Admin{Username: "aaa", Password: "bbb", Avatar: "ccc"})
	admin := &model.Admin{Username: "aaa", Password: "bbb", Avatar: "ccc"}
	global.DB().Insert(admin)
	return ctx.JSONOK()
}

func (UserController) Get1(ctx common.HTTPContext) error {
	id := ctx.ParamInt("id")
	ctx.Put("sss", "sdfsdf")

	global.DB().DeleteByID(&model.Admin{}, id)
	return ctx.JSONOK()
}
