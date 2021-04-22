package kaleid

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func Decrypt(_encSrc, _key string) (string, error) {
	key, _ := hex.DecodeString(_key)
	enc, _ := hex.DecodeString(_encSrc)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("[kaleid|Decrypt] create new cipher fail: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("[kaleid|Decrypt] create new aes GCM fail: %v", err)
	}

	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("[kaleid|Decrypt] decrypt fail: %v", err)
	}

	return fmt.Sprintf("%s", plaintext), nil
}
