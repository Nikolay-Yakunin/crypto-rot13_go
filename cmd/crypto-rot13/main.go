package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Nikolay-Yakunin/crypto-rot13_go/internal/crypto"
	"github.com/Nikolay-Yakunin/crypto-rot13_go/pkg/middleware"
)

// TODO: Prometheus metrics
// TODO: Logging
// TODO: Fix env

func corsInit() cors.Config {
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:3000"}

	// Read CORS_ORIGINS from environment variable
	// and split by comma to allow multiple origins
	// Example: CORS_ORIGINS=http://example.com,http://another-example.com
	origins := os.Getenv("CORS_ORIGINS")
	if origins != "" {
		for _, origin := range strings.Split(origins, ",") {
			origin = strings.TrimSpace(origin)
			if origin != "" {
				config.AllowOrigins = append(config.AllowOrigins, origin)
			}
		}
	}

	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	return config
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Миддлвар для CORS
	r.Use(cors.New(corsInit()))

	// Миддлвар для метрик Prometheus
	middleware.PrometheusInit()
	r.Use(middleware.TrackMetrics())

	// Эндпоинт для проверки запуска
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "crypto-rot13_go is running")
	})

	// Эндпоинт чисто для проверки работы
	r.GET("/status", func(c *gin.Context) {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		c.JSON(http.StatusOK, gin.H{
			"sys time":       memStats.Sys,
			"alloc":          memStats.Alloc,
			"total alloc":    memStats.TotalAlloc,
			"num goroutines": runtime.NumGoroutine(),
			"cpu":            runtime.NumCPU(),
			"go version":     runtime.Version(),
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	apiv1 := r.Group("/api/v1")

	// Crypto handler
	apiv1.POST("/crypto", crypto.CryptoHandler)

	return r
}

func main() {
	// Env
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := setupRouter()

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
