package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sum-project/CookiApp/auth-service/cmd/models"
	"net/http"
)

type registerRequest struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

func (s *Server) register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, errorRequest(errors.New("passwords not equals")))
		return
	}

	role, _ := json.Marshal("[ROLE_USER]")
	user := models.User{
		ID:       uuid.New(),
		Email:    req.Email,
		Password: req.Password,
		Role:     role,
	}

	u, err := s.store.UserRepository.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}
	if u != nil {
		c.JSON(http.StatusBadRequest, errorRequest(errors.New("user exists")))
		return
	}

	id, err := s.store.UserRepository.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}

	token, err := s.store.TokenRepository.GenerateToken(id.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}

	tokenUrl := fmt.Sprintf("127.0.0.1:8080/confirm_user/%s", token)

	err = s.sendMail(MailPayload{
		From:    "register@cookapp.com",
		To:      user.Email,
		Subject: "Potwierdz swoje konto",
		Message: fmt.Sprintf("Potwierdz swoje konto klikajac w ten link: <a href=\"%s\">Link</a>", tokenUrl),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorRequest(err))
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": id,
	})
}

func (s *Server) confirmUser(c *gin.Context) {
	type Token struct {
		Value string `uri:"token" binding:"required"`
	}
	var token Token
	if err := c.ShouldBindUri(&token); err != nil {
		c.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}

	id, err := s.store.TokenRepository.ValidToken(token.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}

	if id == "" {
		c.JSON(http.StatusBadRequest, errorRequest(errors.New("invalid token")))
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}

	user, err := s.store.UserRepository.GetUserByID(uid)

	user.Active = true
	err = s.store.UserRepository.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"id": user.ID,
	})
}
