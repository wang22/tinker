package global

import (
	"github.com/jinzhu/gorm"
	"github.com/wang22/tinker/db"
)

// ================== DataBase ==================

var database *db.DataBase
var databaseSeted = false

func SetDataBase(db *db.DataBase) {
	if databaseSeted {
		return
	}
	database = db
	databaseSeted = true
}

func DB() *gorm.DB {
	return database.Config.DB
}

// ================== /DataBase ==================
