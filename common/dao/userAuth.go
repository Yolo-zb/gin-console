package dao

import (
	localGorm "console/src/gorm"
	"github.com/jinzhu/gorm"
	"console/common/model"
)

type UserAuth struct {
	gorm *gorm.DB
}

func NewUserAuth() UserAuth {
	return UserAuth{
		gorm:localGorm.GetGorm("localhost"),
	}
}

func (ctl *UserAuth) GetById(id int) model.UserAuth {
	userAuth := model.UserAuth{}
	ctl.gorm.Where("uid = ?", id).First(&userAuth)
	return userAuth
}
