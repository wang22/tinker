package service

import (
	"github.com/wang22/tinker/db"
)

var AdminDB *db.AdminDB

func Init() {
	AdminDB = &db.AdminDB{}
}
