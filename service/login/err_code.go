package login

import "github.com/jiaxwu/go_service_req_rsp_errcode_demo/service/common"

var (
	ErrCodeInvalidParameterUsernameOrPassword = common.NewErrCode("InvalidParameterUsernameOrPassword",
		"The username or password is not valid.", "账号或密码错误")
)
