package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const TokenExpireDuration = 24 * time.Hour //过期时间
var hmacSampleSecret = []byte("Secret")

type AuthClaim struct {
	Uid int64 `json:"uid"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT
func GenerateToken(uid int64) (tokenStr string) {
	var authClaim AuthClaim
	authClaim.Uid = uid
	authClaim.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)) // 过期时间
	//authClaim.RegisteredClaims.IssuedAt = jwt.NewNumericDate(time.Now())                           // 签发时间
	//authClaim.RegisteredClaims.NotBefore = jwt.NewNumericDate(time.Now())                          // 生效时间
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaim)
	tokenString, _ := token.SignedString(hmacSampleSecret) //私钥加密
	return tokenString
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (auth AuthClaim, Valid bool) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	Valid = token.Valid //token是否有效 true有效  false无效
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		auth.Uid = int64(claims["uid"].(float64))
		auth.ExpiresAt = jwt.NewNumericDate(time.Unix(int64(claims["uid"].(float64)), 0)) // 过期时间
	}
	return
}
