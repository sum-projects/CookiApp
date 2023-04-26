package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sum-project/CookiApp/gateway-service/cmd/request"
	"io"
	"net/http"
)

type Server struct {
	config Config
	router *gin.Engine
}

func NewServer(config Config) *Server {
	server := &Server{
		config: config,
	}
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/handle", server.handle)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorRequest(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (s *Server) sendLogin(msg request.LoginPayload) error {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	return sendRequest(fmt.Sprintf("%s/%s", s.config.AuthService, "login"), jsonData)
}

func (s *Server) sendRegister(msg request.RegisterPayload) error {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	return sendRequest(fmt.Sprintf("%s/%s", s.config.AuthService, "register"), jsonData)
}

func (s *Server) sendAccountConfirm(msg request.AccountConfirmPayload) error {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	return sendRequest(fmt.Sprintf("%s/%s", s.config.AuthService, "confirm_user"), jsonData)
}

func (s *Server) sendMail(msg request.MailPayload) error {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	return sendRequest(s.config.MailService, jsonData)
}

func sendRequest(addr string, data []byte) error {
	req, err := http.NewRequest("POST", addr, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error calling mail service: %v", err)
		}
		return fmt.Errorf("error calling mail service: %s", resBody)
	}

	return err
}
