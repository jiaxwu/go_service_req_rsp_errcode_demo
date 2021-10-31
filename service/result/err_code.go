package result

import "net/http"

var (
	ErrCodeOK               = newErrCode(http.StatusOK, "OK", "", "")
	ErrCodeInvalidParameter = newErrCode(http.StatusBadRequest, "InvalidParameter",
		"The required parameter is not valid.", "非法参数")
	ErrCodeInvalidParameterUsernameOrPassword = newErrCode(http.StatusBadRequest,
		"InvalidParameter.UsernameOrPassword", "The username or password is not valid.", "账号或密码错误")
	ErrCodeInternalError = newErrCode(http.StatusInternalServerError, "InternalError",
		"The request processing has failed due to some unknown error.", "给您带来的不便，深感抱歉，请稍后再试")
)

// ErrCode 错误码
type ErrCode interface {
	Status() int    // HTTP状态码
	Code() string   // 错误码
	Msg() string    // 错误消息
	Advice() string // 建议处理方式
}

// errCodeImpl 错误码实现
type errCodeImpl struct {
	status int
	code   string
	msg    string
	advice string
}

func (e errCodeImpl) Status() int {
	return e.status
}

func (e errCodeImpl) Code() string {
	return e.code
}

func (e errCodeImpl) Msg() string {
	return e.msg
}

func (e errCodeImpl) Advice() string {
	return e.advice
}

// newErrCode 新建一个错误码
func newErrCode(status int, code, msg, advice string) ErrCode {
	return errCodeImpl{
		code:   code,
		msg:    msg,
		status: status,
		advice: advice,
	}
}
