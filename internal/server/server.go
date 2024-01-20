package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	configs "news/internal/config"
	"news/internal/controller/http"
	"news/internal/service/usecase"
	db "news/pkg/postgres"
	"news/pkg/logger"
)

type Server struct{
	cfg *configs.Config
	log logger.Logger
}
func NewServer(
	cfg *configs.Config,
	log logger.Logger,
	) *Server{
	return &Server{
		cfg: cfg,
		log: log,
	}
}
func (s Server) Run() error{
	pg,err :=db.DB(&s.cfg.Postgres)
	if err != nil {
		s.log.Fatal(err)
	}
	r:=gin.New()
	uc:=usecase.New(pg,s.log)
	http.SetUp(r,uc)
	return r.Run(fmt.Sprintf(":%d", s.cfg.Server.Port))
}