package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sum-project/CookiApp/gateway-service/cmd/request"
	"net/http"
)

func (s *Server) handle(c *gin.Context) {
	var req request.RequestPayload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}

	switch req.Action {
	case "login":
		if err := s.login(req); err != nil {
			c.JSON(http.StatusInternalServerError, errorRequest(err))
		}
	case "register":
		if err := s.register(req); err != nil {
			c.JSON(http.StatusInternalServerError, errorRequest(err))
		}
	case "accountConfirm":
		if err := s.accountConfirm(req); err != nil {
			c.JSON(http.StatusInternalServerError, errorRequest(err))
		}
	case "mail":
		if err := s.mail(req); err != nil {
			c.JSON(http.StatusInternalServerError, errorRequest(err))
		}
	default:
		c.JSON(http.StatusBadRequest, errorRequest(errors.New("unknown action")))
	}
}

func (s *Server) login(req request.RequestPayload) error {
	return s.sendLogin(req.Login)
}

func (s *Server) register(req request.RequestPayload) error {
	return s.sendRegister(req.Register)
}

func (s *Server) accountConfirm(req request.RequestPayload) error {
	return s.sendAccountConfirm(req.AccountConfirm)
}

func (s *Server) mail(req request.RequestPayload) error {
	return s.sendMail(req.Mail)
}
