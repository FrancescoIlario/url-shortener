package idgen

import (
	"crypto/rand"

	"github.com/btcsuite/btcutil/base58"
)

// NewID generates a new string ID
func NewID() (string, error) {
	r := make([]byte, 5)
	if _, err := rand.Read(r); err != nil {
		return "", err
	}

	// use base58 instead of base64 to avoid some special chars like '/+='
	s := base58.Encode(r)
	return s, nil
}
