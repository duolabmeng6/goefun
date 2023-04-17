package os

import "github.com/gogf/gf/crypto/gaes"

// Encrypt is alias of EncryptCBC.
func Aes加密(plainText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	return gaes.EncryptCBC(plainText, key, iv...)
}

// Decrypt is alias of DecryptCBC.
func Aes解密(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	return gaes.DecryptCBC(cipherText, key, iv...)
}

func AesCFB加密(plainText []byte, key []byte, padding *int, iv ...[]byte) ([]byte, error) {
	return gaes.EncryptCFB(plainText, key, padding, iv...)
}

func AesCFB解密(cipherText []byte, key []byte, unPadding int, iv ...[]byte) ([]byte, error) {
	return gaes.DecryptCFB(cipherText, key, unPadding, iv...)

}
