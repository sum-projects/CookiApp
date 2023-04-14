package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type mailMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (s *Server) SendMail(c *gin.Context) {
	var req mailMessage
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}

	msg := Message{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
		Data:    req.Message,
	}

	if err := s.mailer.SendSMTPMessage(msg); err != nil {
		c.JSON(http.StatusInternalServerError, errorRequest(err))
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": fmt.Sprintf("sent to %s", req.To),
	})
}
