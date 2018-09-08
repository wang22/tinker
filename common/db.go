package common

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBConfig struct {
	DB          *gorm.DB
	DBType      string
	DBName      string
	Port        int
	Host        string
	Username    string
	Password    string
	TablePrefix string
	LogEnable   bool
}

type DataBase struct {
	Config *DBConfig
}

func (db *DataBase) Init(dbc *DBConfig) error {
	if dbc.DBType == "mysql" {
		conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbc.Username, dbc.Password, dbc.Host, dbc.Port, dbc.DBName)
		db, err := gorm.Open("mysql", conn)
		if err != nil {
			return err
		}
		dbc.DB = db
	} else if dbc.DBType == "sqlite" {
		// TODO sqlite
	} else {
		return errors.New("Unsupport database type:" + dbc.DBType)
	}
	dbc.DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbc.TablePrefix + defaultTableName
	}
	dbc.DB.LogMode(dbc.LogEnable)
	db.Config = dbc
	return nil
}

func (db *DataBase) Insert(model interface{}) {
	t := reflect.ValueOf(model)
	method := t.MethodByName("InsertHook")
	if method.IsValid() {
		method.Call(nil)
	}
	db.Config.DB.Save(model)
}

func (db *DataBase) Update(model interface{}) {
	t := reflect.ValueOf(model)
	method := t.MethodByName("UpdateHook")
	if method.IsValid() {
		method.Call(nil)
	}
	db.Config.DB.Model(model).Update(model)
}

func (db *DataBase) Delete(model interface{}) {
	t := reflect.ValueOf(model)
	method := t.MethodByName("DeleteHook")
	if method.IsValid() {
		method.Call(nil)
	}
	db.Config.DB.Model(model).Update(model)
}

func (db *DataBase) DeleteByID(model interface{}, ID int) {
	db.FetchByID(model, ID)
	db.Delete(model)
}

func (db *DataBase) FetchByID(model interface{}, ID int) {
	db.Config.DB.First(model, ID)
}

// ================== DataBase ==================
var database *DataBase
var databaseSeted = false

func SetDataBase(db *DataBase) {
	if databaseSeted {
		return
	}
	database = db
	databaseSeted = true
}

func DB() *DataBase {
	return database
}

func GormDB() *gorm.DB {
	return database.Config.DB
}

// ================== /DataBase ==================
