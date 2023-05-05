package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"time"

	"github.com/gin-contrib/cors"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func StartRoute() {
	mapUrls()

	log.Info("Starting server")
	router.Run(":8090")

}
