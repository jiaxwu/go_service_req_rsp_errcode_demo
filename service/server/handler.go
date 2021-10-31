package server

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohuashifu/go_service_req_rsp_errcode_demo/service/result"
	"reflect"
)

func WrapService(service interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		// 参数绑定
		s := reflect.TypeOf(service)
		reqPointType := s.In(0)
		reqStructType := reqPointType.Elem()
		req := reflect.New(reqStructType)
		if err := c.ShouldBindJSON(req.Interface()); err != nil {
			result.Failure(c, result.ErrCodeInvalidParameter)
			return
		}

		// 调用服务
		params := []reflect.Value{reflect.ValueOf(req.Interface())}
		rets := reflect.ValueOf(service).Call(params)

		// 结果处理
		if !rets[1].IsNil() {
			result.Failure(c, rets[1].Interface().(result.Error))
			return
		}
		result.Success(c, rets[0].Interface())
	}
}
