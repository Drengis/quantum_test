package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/quantum-server/internal/dto"
	"github.com/user/quantum-server/internal/service"
)

// @title Quantum Mortgage API
// @version 1.0
// @description API для расчёта ипотечных профилей
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /

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

// @Summary Создать ипотечный профиль
// @Description Создаёт новый расчёт ипотеки и возвращает ID задачи
// @Tags Mortgage
// @Accept json
// @Produce json
// @Param request body dto.CreateMortgageRequest true "Параметры ипотеки"
// @Success 202 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /mortgage-profiles [post]
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

// @Summary Получить расчёт ипотеки
// @Description Возвращает расчёт ипотеки по ID
// @Tags Mortgage
// @Accept json
// @Produce json
// @Param id path int true "ID расчёта"
// @Success 200 {object} dto.MortgageResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /mortgage-profiles/{id} [get]
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
