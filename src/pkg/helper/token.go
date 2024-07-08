package helper

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"simple-api.com/m/src/config"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
)

type Claims struct {
	Id       int64  `json:"id"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Status   string `json:"status,omitempty"`
}

type AccessClaims struct {
	Claims
	jwt.StandardClaims
}

func GenerateAccessToken(ctx context.Context, claims Claims) (string, error) {
	log := logger.NewLogger()
	cfg, _ := config.LoadEnv(ctx, log)
	jwtClaims := AccessClaims{Claims: claims, StandardClaims: jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(cfg.Jwt.AccessTokenExpired) * time.Hour).Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	accessToken, err := token.SignedString([]byte(cfg.Jwt.SecretKey))
	if err != nil {
		log.Error(ctx, err.Error(), err, nil)
		return "", err
	}

	return accessToken, nil
}

func VerifyToken(ctx context.Context, tokenString string) (*AccessClaims, error) {
	log := logger.NewLogger()
	cfg, _ := config.LoadEnv(ctx, log)
	secretKey := []byte(cfg.Jwt.SecretKey)
	token, err := jwt.ParseWithClaims(tokenString, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Error(ctx, err.Error(), err, nil)
		return nil, err
	}

	claims, ok := token.Claims.(*AccessClaims)
	if !ok || !token.Valid {
		log.Error(ctx, "Error token not valid", err, nil)
		return nil, wrapper.UnauthorizedError("invalid token")
	}

	return claims, nil
}
