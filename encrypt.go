package kaleid

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func Encrypt(_src, _key string) (string, error) {
	bytes := []byte(_src)
	key, _ := hex.DecodeString(_key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("[kaleid|encrypt] create new cipher fail: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("[kaleid|encrypt] create new aes GCM fail: %v", err)
	}

	nonce := make([]byte, aesGCM.NonceSize())

	cipherText := aesGCM.Seal(nonce, nonce, bytes, nil)
	return fmt.Sprintf("%x", cipherText), nil
}
