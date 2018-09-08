package controller_admin

import (
	"net/http"
	"time"

	"github.com/wang22/tinker/common"
	"github.com/wang22/tinker/config"
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
	admin := service.Admin.FindByUsername(adminLoginForm.Username)
	if !self.comparisonPWD(admin.Password, adminLoginForm.Password) {
		return ctx.JSONErr("用户名或密码错误")

	}
	cookie := new(http.Cookie)
	cookie.Name = config.Get(config.ConstAuthCookieKey)
	cookie.Value = "22"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	ctx.SetCookie(cookie)

	ctx.Put("admin", admin)
	return ctx.JSONOK()
}

func (AdminController) comparisonPWD(oldPwd string, newPwd string) bool {
	return oldPwd == newPwd
}
