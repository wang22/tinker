package main

import (
	"github.com/labstack/echo"
	"github.com/wang22/tinker/app"
	"github.com/wang22/tinker/config"
	"github.com/wang22/tinker/controller/admin"
)

func main() {
	app.Initial()

	e := echo.New()
	adminRouter := controller_admin.AdminRouter{}
	adminRouter.Routing(e.Group(config.Get(config.ConstHttpAdminPrefix)))

	e.Logger.Fatal(e.Start(config.Get(config.ConstHttpPort)))
}
