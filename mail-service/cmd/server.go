package main

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	mailer Mail
	router *gin.Engine
}

func NewServer(mailer Mail) *Server {
	server := &Server{
		mailer: mailer,
	}
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.router = router
	return server
}

func (s Server) Start(address string) error {
	return s.router.Run(address)
}

func errorRequest(err error) gin.H {
	return gin.H{"error": err.Error()}
}
