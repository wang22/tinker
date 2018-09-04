package main

import (
	"github.com/wang22/tinker/config"
		"github.com/labstack/echo"
		"github.com/wang22/tinker/controller/admin"
)

func main() {

	config.Initial("config.yml")
	e := echo.New()

	adminRouter := controller_admin.AdminRouter{}
	adminRouter.Routing(e.Group(config.Get(config.CONST_HTTP_ADMIN_PREFIX)))

	e.Logger.Fatal(e.Start(config.Get(config.CONST_HTTP_PORT)))

}
