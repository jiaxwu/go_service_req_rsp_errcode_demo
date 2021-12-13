package login

import (
	"github.com/google/uuid"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Login(req *LoginReq) (*LoginRsp, error) {
	if req.Username != "admin" || req.Password != "admin" {
		return nil, ErrCodeInvalidParameterUsernameOrPassword
	}

	return &LoginRsp{
		Token: uuid.New().String(),
	}, nil
}
