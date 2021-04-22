package kaleid

import (
	"golang.org/x/crypto/sha3"
)

func GenerateKey(_src string) []byte {
	h := make([]byte, 32)
	sha3.ShakeSum256(h, []byte(_src))

	return h
}
