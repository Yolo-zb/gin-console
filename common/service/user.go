package service

import (
	"console/common/dao"
	"console/common/model"
)

// 一般情况下service应该只包含并调用自己的data模型，需要其他服务的功能请service.Xxx调用服务而不是引入其他data模型
var User = userService{
	data: dao.NewUser(),
}

type userService struct {
	data dao.User
}

func (ctl *userService) GetById(id int) model.User {
	return ctl.data.GetById(id)
}
