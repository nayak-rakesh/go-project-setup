package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nayak-rakesh/go-project-setup/domain"
)

// UserHandler ...
type UserHandler struct {
	userService domain.UserSerice
}

// NewUserHandler ...
func NewUserHandler(router *gin.Engine, service domain.UserSerice) {
	handler := &UserHandler{
		userService: service,
	}
	router.POST("user/create", handler.Create)
}

// Create ...
func (h *UserHandler) Create(c *gin.Context) {
	var user domain.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = h.userService.Create(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, user)
}
