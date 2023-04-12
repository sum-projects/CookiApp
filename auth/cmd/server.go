package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sum-project/CookiApp/auth/cmd/db"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
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
