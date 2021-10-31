package server

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohuashifu/go_service_req_rsp_errcode_demo/service/login/service"
)

func initRouter(r *gin.Engine) {
	var loginService service.LoginService
	r.POST("login", WrapService(loginService.Login))
}

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	initRouter(r)
	return &Server{
		router: r,
	}
}

func (s *Server) Run() error {
	return s.router.Run()
}
