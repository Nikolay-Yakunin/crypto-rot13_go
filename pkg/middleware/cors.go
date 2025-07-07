package middleware

import (
	"github.com/gin-contrib/cors"

	"time"
)

type Cors struct {
	Origins []string
}

func (c Cors)CorsInit() cors.Config {
	config := cors.DefaultConfig()

	config.AllowOrigins = c.Origins

	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	return config
}
