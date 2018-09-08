package service

import (
	"github.com/wang22/tinker/common"
	"github.com/wang22/tinker/model"
)

// =========== Init service ===========
var Admin *AdminService

func Init() {
	Admin = &AdminService{}
}

// =========== Init service ===========

type AdminService struct{}

func (AdminService) FindByUsername(username string) *model.Admin {
	admin := &model.Admin{}
	common.GormDB().Where("username = ?", username).First(admin)
	return admin
}
