package result

import (
	"github.com/gin-gonic/gin"
)

// Rsp 响应
type Rsp struct {
	Code string      `json:"code,omitempty"` // 错误码
	Msg  string      `json:"msg,omitempty"`  // 消息
	Data interface{} `json:"data,omitempty"` // 数据
}

// Success 请求成功
func Success(c *gin.Context, data interface{}) {
	rsp(c, ErrCodeOK, data)
}

// Failure 请求失败
func Failure(c *gin.Context, errCode ErrCode) {
	rsp(c, errCode, nil)
}

// rsp 响应
func rsp(c *gin.Context, errCode ErrCode, data interface{}) {
	c.JSON(errCode.Status(), &Rsp{
		Code: errCode.Code(),
		Msg:  errCode.Advice(),
		Data: data,
	})
}
