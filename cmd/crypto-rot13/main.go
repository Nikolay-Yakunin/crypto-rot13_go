package main

import (
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

// TODO: CORS
// TODO: Prometheus metrics
// TODO: Logging

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Тут бы корс добавить

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "crypto-rot13_go is running")
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Header("Content-Type", "image/x-icon")
		c.File("../../static/favicon.ico")
	})

	r.GET("/status", func(c *gin.Context) {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		c.JSON(http.StatusOK, gin.H{
			"sys time":   memStats.Sys,
			"alloc":      memStats.Alloc,
			"total alloc": memStats.TotalAlloc,
			"num goroutines": runtime.NumGoroutine(),
			"cpu":        runtime.NumCPU(),
			"go version": runtime.Version(),
		})
	})

	// r.GET("prometheus", func(c *gin.Context) {
		
	// }

	return r;
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := setupRouter()
	r.Run(":" + port)
}
