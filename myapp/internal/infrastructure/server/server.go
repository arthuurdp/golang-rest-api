package server

import (
	"log"

	"github.com/gin-gonic/gin"

	"Hello_World/myapp/internal/adapter/handler"
)

type Server struct {
	router *gin.Engine
	userHandler *handler.UserHandler
}

func NewServer(userHandler *handler.UserHandler) *Server {
	return &Server{
		router: gin.Default(),
		userHandler: userHandler,
	}
}

func (s *Server) Setup() {
	s.router.Use(gin.Recovery())
	s.router.Use(gin.Logger())

	v1 := s.router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", s.userHandler.FindAll)
			users.GET("/:id", s.userHandler.FindById)
			users.POST("", s.userHandler.Create)
			users.PATCH("/:id", s.userHandler.Update)
			
		}
	}
}

func (s *Server) Run(address string) error {
	log.Printf("Server running at %s", address)
	return s.router.Run(address)
}