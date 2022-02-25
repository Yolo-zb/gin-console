package dao

import (
	"console/common/model"
	localGorm "console/src/gorm"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm localGorm.A
	b *gorm.DB
}

func NewUser() User {
	return User{
		gorm:localGorm.A{},
		b:localGorm.GetGorm("localhost"),
	}
}

func (ctl *User) GetById(id int) model.User {
	user := model.User{}
	ctl.gorm.GetGorm("localhost").Where("id = ?", id).First(&user)
	return user
}
