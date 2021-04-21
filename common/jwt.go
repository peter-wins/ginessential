package common

import (
	"ginessential/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)
// jwt认证
var jwtKey = []byte("a_secret_crect")

type Claims struct{
	UserId uint
	jwt.StandardClaims
}
func ReleaseToken(user model.User) (string, error){
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // TOKEN过期时间
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "oceanlearn.tech",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString,nil
}
// ParseToken 从tokenString里面解析出claims，然后返回
func ParseToken(tokenString string)(*jwt.Token, *Claims, error){
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey,  nil
	})

	return token, claims, err
}