package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/user/quantum-server/config"
	db "github.com/user/quantum-server/internal/database"
	"github.com/user/quantum-server/internal/dto"
	"github.com/user/quantum-server/internal/repository"
	"github.com/user/quantum-server/internal/service"
	"github.com/user/quantum-server/internal/worker"
)

func main() {
	cfg := config.LoadConfig()
	dbConn := db.Init(cfg.DB)

	profileRepo := repository.NewMortgageProfileRepository(dbConn)
	calcRepo := repository.NewMortgageCalculationRepository(dbConn)

	// Channel
	taskChan := make(chan dto.MortgageTask, 100)

	// Service
	mortgageService := service.NewMortgageService(dbConn, profileRepo, calcRepo, taskChan)

	// Worker
	calcWorker := worker.NewCalculationWorker(mortgageService, calcRepo, taskChan)
	calcWorker.Start(context.Background())

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

	r.Run(cfg.AppURL + ":" + cfg.MainAppPort)
}
