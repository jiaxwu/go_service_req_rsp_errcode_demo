package service

import (
	"github.com/google/uuid"
	"github.com/xiaohuashifu/go_service_req_rsp_errcode_demo/service/login/model"
	"github.com/xiaohuashifu/go_service_req_rsp_errcode_demo/service/result"
)

type LoginService struct{}

func (s *LoginService) Login(req *model.LoginReq) (*model.LoginRsp, result.Error) {
	if req.Username != "admin" || req.Password != "admin" {
		return nil, result.NewError(result.ErrCodeInvalidParameterUsernameOrPassword)
	}

	return &model.LoginRsp{
		Token: uuid.New().String(),
	}, nil
}
