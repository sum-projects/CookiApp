package main

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sum-project/CookiApp/auth/cmd/models"
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

	c.JSON(http.StatusAccepted, gin.H{
		"message": id,
	})
}
