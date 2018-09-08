package app

import (
	"github.com/coocood/freecache"
	"github.com/wang22/tinker/common"
	"github.com/wang22/tinker/config"
	"github.com/wang22/tinker/service"
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

func initDataBase() {
	database := &common.DataBase{}
	err := database.Init(createDBConfig())
	if err != nil {
		panic(err)
	}
	common.SetDataBase(database)
}

func initCache() {
	cacheType := config.Get(config.ConstCacheType)
	var cacher common.Cacher
	if cacheType == "redis" {
		redis := &common.RedisCacher{}
		cacher = redis
	} else if cacheType == "gocache" {
		gocache := &common.GOCacher{}
		gocache.Cache = freecache.NewCache(config.GetInt(config.ConstCacheSize))
		cacher = gocache
	} else {
		panic("Unsupport cache type:" + cacheType)
	}
	common.Cache = cacher
}

func Initial() {
	config.Initial("config.yml")
	initDataBase()
	initCache()
	service.Init()
}
