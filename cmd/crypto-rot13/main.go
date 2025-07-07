package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Nikolay-Yakunin/crypto-rot13_go/internal/crypto"
	"github.com/Nikolay-Yakunin/crypto-rot13_go/pkg/middleware"
)

func setupMiddleware(r *gin.Engine) {
	c := middleware.Cors{
		Origins: strings.Split(os.Getenv("CORS_ORIGINS"), ","),
	}
	r.Use(cors.New(c.CorsInit()))

	middleware.PrometheusInit() // Я просто скопировал этот миддлвар, так что он не на типах
	r.Use(middleware.TrackMetrics())

}

func setupRouter() *gin.Engine {
	r := gin.Default()

	setupMiddleware(r)

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
