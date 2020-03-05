package os

import (
	"github.com/gogf/gf/crypto/gdes"
)

// EncryptECB encrypts <plainText> using ECB mode.
func DesECB加密(plainText []byte, key []byte, padding int) ([]byte, error) {
	return gdes.EncryptECB(plainText, key, padding)
}

// EncryptECB encrypts <plainText> using ECB mode.
func DesECB解密(cipherText []byte, key []byte, padding int) ([]byte, error) {
	return gdes.DecryptECB(cipherText, key, padding)
}
