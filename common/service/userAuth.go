package service

import (
	"github.com/Yolo-zb/gin-console/common/dao"
	"github.com/Yolo-zb/gin-console/common/model"
)

// 一般情况下service应该只包含并调用自己的data模型，需要其他服务的功能请service.Xxx调用服务而不是引入其他data模型
var UserAuth = userAuthService{
	data: dao.NewUserAuth(),
}

type userAuthService struct {
	data dao.UserAuth
}

func (ctl *userAuthService) GetById(id int) model.UserAuth {
	return ctl.data.GetById(id)
}
