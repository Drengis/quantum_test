package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/quantum-server/internal/dto"
	"github.com/user/quantum-server/internal/service"
)

type MortgageHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
}

type mortgageHandler struct {
	mortgageService service.MortgageService
}

func NewMortgageHandler(mortgageService service.MortgageService) MortgageHandler {
	return &mortgageHandler{mortgageService: mortgageService}
}

func (h *mortgageHandler) Create(c *gin.Context) {
	var req dto.CreateMortgageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	calc, err := h.mortgageService.CreateCalculation(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"id":     calc.ID,
		"status": calc.Status,
	})
}

func (h *mortgageHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	resp, err := h.mortgageService.GetCalculation(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resp == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Calculation not found"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
