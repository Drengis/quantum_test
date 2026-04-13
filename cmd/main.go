package main

import (
	"github.com/gin-gonic/gin"
	"github.com/user/quantum-server/config"
	db "github.com/user/quantum-server/internal/database"
)

func main() {
	cfg := config.LoadConfig()
	db.Init(cfg.DB)

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
