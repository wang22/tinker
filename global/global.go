package global

import (
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

func DB() *db.DataBase {
	return database
}

// ================== /DataBase ==================
