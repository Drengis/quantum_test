package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/user/quantum-server/config"
	db "github.com/user/quantum-server/internal/database"
	"github.com/user/quantum-server/internal/dto"
	"github.com/user/quantum-server/internal/handler"
	"github.com/user/quantum-server/internal/repository"
	"github.com/user/quantum-server/internal/service"
	"github.com/user/quantum-server/internal/worker"

	_ "github.com/user/quantum-server/docs"
)

// @title Quantum Mortgage API
// @version 1.0
// @description API для расчёта ипотечных профилей
// @host localhost:8081
// @BasePath /
func main() {
	cfg := config.LoadConfig()
	dbConn := db.Init(cfg.DB)
	rdb := db.InitRedis(cfg.Redis)

	profileRepo := repository.NewMortgageProfileRepository(dbConn)
	calcRepo := repository.NewMortgageCalculationRepository(dbConn)
	userRepo := repository.NewUserRepository(dbConn)

	taskChan := make(chan dto.MortgageTask, 100)

	mortgageService := service.NewMortgageService(dbConn, profileRepo, calcRepo, rdb, taskChan)
	userService := service.NewUserService(userRepo)

	calcWorker := worker.NewCalculationWorker(mortgageService, calcRepo, taskChan)
	calcWorker.Start(context.Background())

	mortgageHandler := handler.NewMortgageHandler(mortgageService)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

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

	r.POST("/mortgage-profiles", mortgageHandler.Create)
	r.GET("/mortgage-profiles/:id", mortgageHandler.Get)

	r.POST("/user", userHandler.FindOrCreate)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})

	r.Run(":" + cfg.MainAppPort)
}
