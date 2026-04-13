package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/user/quantum-server/config"
	db "github.com/user/quantum-server/internal/database"
	"github.com/user/quantum-server/internal/dto"
	"github.com/user/quantum-server/internal/handler"
	"github.com/user/quantum-server/internal/repository"
	"github.com/user/quantum-server/internal/service"
	"github.com/user/quantum-server/internal/worker"
)

func main() {
	cfg := config.LoadConfig()
	dbConn := db.Init(cfg.DB)
	rdb := db.InitRedis(cfg.Redis)

	profileRepo := repository.NewMortgageProfileRepository(dbConn)
	calcRepo := repository.NewMortgageCalculationRepository(dbConn)

	// Channel
	taskChan := make(chan dto.MortgageTask, 100)

	// Service
	mortgageService := service.NewMortgageService(dbConn, profileRepo, calcRepo, rdb, taskChan)

	// Worker
	calcWorker := worker.NewCalculationWorker(mortgageService, calcRepo, taskChan)
	calcWorker.Start(context.Background())

	// Handlers
	mortgageHandler := handler.NewMortgageHandler(mortgageService)

	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Routes
	r.POST("/mortgage-profiles", mortgageHandler.Create)
	r.GET("/mortgage-profiles/:id", mortgageHandler.Get)

	r.Run(":" + cfg.MainAppPort)
}
