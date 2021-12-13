package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(engine *gin.Engine) *Server {
	return &Server{
		engine: engine,
	}
}

func (s *Server) Run() error {
	return s.engine.Run()
}
