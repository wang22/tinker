package db

import (
	"github.com/wang22/tinker/common"
	"github.com/wang22/tinker/model"
)

type AdminDB struct{}

func (AdminDB) FindByUsername(username string) *model.Admin {
	admin := &model.Admin{}
	common.GormDB().Where("username = ?", username).First(admin)
	return admin
}
