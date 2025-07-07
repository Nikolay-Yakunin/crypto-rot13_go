# crypto-rot13_go

[![Go Report Card](https://goreportcard.com/badge/github.com/Nikolay-Yakunin/crypto-rot13_go)](https://goreportcard.com/report/github.com/Nikolay-Yakunin/crypto-rot13_go)

REST API на Go для шифрования и дешифровки текста методом [ROT13](https://ru.wikipedia.org/wiki/ROT13) и другими методами шифрования.  
---
## Endpoints

### POST /api/v1/crypto
**method** - можно получить по /api/v1/crypto/methods
**mode** - **'encrypt'** для расшифровки, **'decrypt'** для шифрования.

```go
Method string `json:"method" binding:"required"`
Mode   string `json:"mode" binding:"required,oneof=encrypt decrypt"`
```
Пример:
```
POST api/v1/crypt?method=rot13&mode=encrypt
GET api/v1/crypt/methods
```

# License

MIT License
