package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jiaxwu/go_service_req_rsp_errcode_demo/service/login"
	"github.com/jiaxwu/go_service_req_rsp_errcode_demo/service/server"
	"log"
)

func main() {
	engine := gin.Default()
	loginService := login.NewService()
	login.RegisterHandler(engine, loginService)
	webServer := server.NewServer(engine)
	log.Println(webServer.Run())
}
