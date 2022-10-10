package middleware

import (
	"App/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("Opus_secret")


// ReleaseToken 生成Token
func ReleaseToken(user model.User) (string, error) {

	// 超时时间
	expirationTime := time.Now().Add(1 * 24 * time.Hour)

	// token结构生成
	claims := &model.Claims{
		// 使用ID作为	有效载荷
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "fzf404.top",
			Subject:   "user token",
		},
	}

	// 编码算法SigningMethodHS256
	// 将Claims加密存储为Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 加入key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解码Token
func ParseToken(tokenString string) (*jwt.Token, *model.Claims, error) {
	claims := &model.Claims{}
	var token *jwt.Token
	var err error
	// 解码
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
