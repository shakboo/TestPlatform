package util

import (
	// "time"
	jwt "github.com/dgrijalva/jwt-go"
	"testplatform/pkg/setting"
)

var(
	jwtSecret = []byte(setting.JwtSecret)
	// expireTime = time.Now().Add(1 * time.Hour)
)
	

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(username, password string) (string, error) {

	claims := Claims{
		username,
		password,
		jwt.StandardClaims {
			// ExpiresAt : expireTime.Unix(),
			Issuer : "testplatform",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
