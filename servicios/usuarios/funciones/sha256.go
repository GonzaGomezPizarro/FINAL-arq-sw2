package funciones

import (
	"crypto/sha256"
	"fmt"
)

// NewSHA256 ...
func SSHA256(data string) string {
	X := []byte(data)
	hash := sha256.Sum256(X)
	a := fmt.Sprintf("%x", hash)
	return a
}
