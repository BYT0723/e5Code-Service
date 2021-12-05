package common

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// return a uuid string
func GetUUID() string {
	return uuid.New().String()
}

func EncryptPwd(password string) (string, error) {
	resultBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}

func ComparePwd(encryptedPwd, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPwd), []byte(pwd)); err != nil {
		return false
	}
	return true
}

func GenerateToken(secretKey string, iat, seconds int64, info map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range info {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
