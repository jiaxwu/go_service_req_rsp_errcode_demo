package login

import (
	"github.com/gin-gonic/gin"
	"github.com/jiaxwu/go_service_req_rsp_errcode_demo/service/wrap"
)

// RegisterHandler 注册handler
func RegisterHandler(engine *gin.Engine, service *Service) {
	engine.POST("/login", wrap.WrapService(service.Login))
}