package controller_admin

import (
	"strconv"

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
	tmp := ctx.Context.Param("id")
	ctx.Put("sss", "sdfsdf")

	id, _ := strconv.Atoi(tmp)
	global.DB().DeleteByID(&model.Admin{}, id)
	return ctx.JSONOK()
}
