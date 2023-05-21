package jwt

import (
	"fmt"
	"testing"
)

func Test_jwttoken(t *testing.T) {
	token := GenerateToken(1)
	fmt.Println(token)

	auth, Valid := ParseToken(token)
	fmt.Println(auth.Uid)
	fmt.Println(Valid)

}
