package common

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
)

// return a uuid string
func GenUUID() string {
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

func GetUserID(ctx context.Context) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return md.Get(UserID)[0], nil
	}
	return "", errors.New("fail to get userid from incomingContext")
}
