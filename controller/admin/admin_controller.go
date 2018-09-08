package controller_admin

import (
	"github.com/wang22/tinker/common"
	"github.com/wang22/tinker/model"
	"github.com/wang22/tinker/service"
)

type AdminController struct{}

func (self *AdminController) Login(ctx common.HTTPContext) error {
	adminLoginForm := new(model.AdminLoginForm)
	err := ctx.Bind(adminLoginForm)
	if err != nil {
		panic(err)
	}
	admin := service.AdminDB.FindByUsername(adminLoginForm.Username)
	if self.comparisonPWD(admin.Password, adminLoginForm.Password) {
		ctx.Put("admin", admin)
		return ctx.JSONOK()
	}
	return ctx.JSONErr("用户名或密码错误")
}

func (AdminController) comparisonPWD(oldPwd string, newPwd string) bool {
	return oldPwd == newPwd
}
