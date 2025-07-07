package crypto

type Crypto struct {
	Method string `json:"method" binding:"required"`
	Mode   string `json:"mode" binding:"required,oneof=encrypt decrypt"`
	Text   string `json:"text" binding:"required"`
}

type Service interface {
	EncryptText() (string, error)
	DecryptText() (string, error)
}

var validMethods = []string{"rot13"}
