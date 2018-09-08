package main

import (
	"github.com/labstack/echo"
	"github.com/wang22/tinker/common"
	"github.com/wang22/tinker/config"
	"github.com/wang22/tinker/controller/admin"
	"github.com/wang22/tinker/global"
)

func createDBConfig() *common.DBConfig {
	dbc := &common.DBConfig{}
	dbc.DBType = config.Get(config.ConstDBType)
	dbc.DBName = config.Get(config.ConstDBName)
	dbc.Port = config.GetInt(config.ConstDBPort)
	dbc.Host = config.Get(config.ConstDBHost)
	dbc.Username = config.Get(config.ConstDBUsername)
	dbc.Password = config.Get(config.ConstDBPassword)
	dbc.TablePrefix = config.Get(config.ConstDBTablePrefix)
	dbc.LogEnable = config.GetBool(config.ConstDBLogEnable)
	return dbc
}

func initial() {
	config.Initial("config.yml")
	database := &common.DataBase{}
	err := database.Init(createDBConfig())
	if err != nil {
		panic(err)
	}
	common.SetDataBase(database)
	global.Init()
}

func main() {
	initial()
	e := echo.New()
	adminRouter := controller_admin.AdminRouter{}
	adminRouter.Routing(e.Group(config.Get(config.ConstHttpAdminPrefix)))
	e.Logger.Fatal(e.Start(config.Get(config.ConstHttpPort)))
}
