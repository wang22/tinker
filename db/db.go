package db

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBConfig struct {
	DB       *gorm.DB
	DBType   string
	DBName   string
	Port     int
	Host     string
	Username string
	Password string
}

type DataBase struct{}

func (DataBase) Init(dbc DBConfig) error {
	if dbc.DBType == "mysql" {
		conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbc.Username, dbc.Password, dbc.Host, dbc.Port, dbc.DBName)
		fmt.Println(conn)
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
	return nil
}
