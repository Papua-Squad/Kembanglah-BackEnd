package app

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kembanglah/config"
)

type Server struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Echo:   echo.New(),
		DB:     NewDatabase(config),
		Config: config,
	}
}

func (server *Server) Start(addr string) error {
	return server.Echo.Start(":" + addr)
}
