package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sum-project/CookiApp/auth-service/cmd/db"
)

type Server struct {
	config Config
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store, config Config) *Server {
	server := &Server{
		store:  store,
		config: config,
	}
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/register", server.register)

	server.router = router
	return server
}

func (s Server) Start(address string) error {
	return s.router.Run(address)
}

func errorRequest(err error) gin.H {
	return gin.H{"error": err.Error()}
}
