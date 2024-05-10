package jwtx

import (
	"github.com/golang-jwt/jwt"
	"net/http"
)

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// 在请求头AUTHORIZATION中获取token
func GetTokenFromHeader(r *http.Request) string {
	bear := r.Header.Get("AUTHORIZATION")
	// 去除前缀后的字符串
	bear = bear[len("Bearer "):]
	return bear
}

// 获取token中的uid
func GetUidFromToken(secretKey string, r *http.Request) (int64, error) {
	token := GetTokenFromHeader(r)
	return getUid(secretKey, token)
}

// 通过token获取uid
func getUid(secretKey, token string) (int64, error) {
	claims := make(jwt.MapClaims)
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	uid := claims["uid"].(float64)
	return int64(uid), nil
}
