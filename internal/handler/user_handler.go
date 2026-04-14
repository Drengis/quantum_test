package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/quantum-server/internal/dto"
	"github.com/user/quantum-server/internal/service"
)

type UserHandler struct {
	service service.UserService
}

type UserHandlerInterface interface {
	FindOrCreate(c *gin.Context)
}

func NewUserHandler(service service.UserService) UserHandlerInterface {
	return &UserHandler{service: service}
}

// @Summary FindOrCreate user
// @Description FindOrCreate user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User data"
// @Success 201 {object} github_com_user_quantum-server_internal_domain.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user [post]
func (h *UserHandler) FindOrCreate(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.FindOrCreate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
