package crypto

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CryptoHandler(r *gin.RouterGroup) {
	r.Use(func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1<<20) // 1 МБ
		c.Next()
	})

	r.GET("/crypt/methods", MethodHandler)
	r.POST("/crypt", CryptHandler)
}

func MethodHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"methods": validMethods,
		"modes":   []string{"encrypt", "decrypt"},
	})
}

func CryptHandler(c *gin.Context) {
	method := c.Query("method")
	if method == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "method is required"})
		return
	}

	mode := c.Query("mode")
	if mode != "encrypt" && mode != "decrypt" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "mode must be 'encrypt' or 'decrypt'"})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot read body"})
		return
	}
	text := string(body)

	result, err := CryptoService(&Crypto{
		Method: method,
		Mode:   mode,
		Text:   text,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
