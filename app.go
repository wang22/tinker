package main

import (
	"github.com/labstack/echo"
	"github.com/wang22/tinker/config"
	"github.com/wang22/tinker/controller/admin"
	"github.com/wang22/tinker/db"
	"github.com/wang22/tinker/global"
)

func createDBConfig() *db.DBConfig {
	dbc := &db.DBConfig{}
	dbc.DBType = config.Get(config.ConstDBType)
	dbc.DBName = config.Get(config.ConstDBName)
	dbc.Port = config.GetInt(config.ConstDBPort)
	dbc.Host = config.Get(config.ConstDBHost)
	dbc.Username = config.Get(config.ConstDBUsername)
	dbc.Password = config.Get(config.ConstDBPassword)
	dbc.TablePrefix = config.Get(config.ConstDBTablePrefix)
	return dbc
}

func initial() {
	config.Initial("config.yml")
	database := &db.DataBase{}
	err := database.Init(createDBConfig())
	if err != nil {
		panic(err)
	}
	global.SetDataBase(database)
}

func main() {
	initial()
	e := echo.New()
	adminRouter := controller_admin.AdminRouter{}
	adminRouter.Routing(e.Group(config.Get(config.ConstHttpAdminPrefix)))
	e.Logger.Fatal(e.Start(config.Get(config.ConstHttpPort)))
}
