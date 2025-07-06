package crypto

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CryptoHandler(c *gin.Context) {
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
