package crypto

import (
	"fmt"
)

func CryptoService(c *Crypto) (string, error) {
	if c == nil {
		return "", fmt.Errorf("crypto service: nil Crypto struct")
	}

	switch c.Mode {
	case "encrypt":
		return c.encryptText()
	case "decrypt":
		return c.decryptText()
	default:
		return "", fmt.Errorf("invalid mode: %s, expected 'encrypt' or 'decrypt'", c.Mode)
	}
}

func (c *Crypto) encryptText() (string, error) {
	if c.Mode != "encrypt" {
		return "", fmt.Errorf("invalid mode: %s, expected 'encrypt'", c.Mode)
	}

	switch c.Method {
	case "rot13":
		return CryptRot13(c.Text), nil
	default:
		return "", fmt.Errorf("unsupported method: %s", c.Method)
	}
}

func (c *Crypto) decryptText() (string, error) {
	if c.Mode != "decrypt" {
		return "", fmt.Errorf("invalid mode: %s, expected 'decrypt'", c.Mode)
	}

	switch c.Method {
	case "rot13":
		return CryptRot13(c.Text), nil // ROT13 шифрование симметрично
	default:
		return "", fmt.Errorf("unsupported method: %s", c.Method)
	}
}
