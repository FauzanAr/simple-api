package helper

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"simple-api.com/m/src/config"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
)

type Role string

const (
	User  Role = "USER"
	Admin Role = "ADMIN"
)

type Claims struct {
	Id       int64  `json:"id"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Status   string `json:"status,omitempty"`
	Role     string `json:"role,omitempty"`
}

type AccessClaims struct {
	Claims
	jwt.StandardClaims
}

func GenerateAccessToken(ctx context.Context, claims Claims, userType Role) (string, error) {
	log := logger.NewLogger()
	cfg, _ := config.LoadEnv(ctx, log)

	var secretKey []byte

	switch userType {
	case User:
		secretKey = []byte(cfg.Jwt.SecretKey)
	case Admin:
		secretKey = []byte(cfg.Jwt.AdminSecretKey)
	default:
		return "", wrapper.InternalServerError("No matching role!")
	}

	jwtClaims := AccessClaims{Claims: claims, StandardClaims: jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(cfg.Jwt.AccessTokenExpired) * time.Hour).Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Error(ctx, err.Error(), err, nil)
		return "", err
	}

	return accessToken, nil
}

func VerifyToken(ctx context.Context, tokenString string, userType Role) (*AccessClaims, error) {
	log := logger.NewLogger()
	cfg, _ := config.LoadEnv(ctx, log)

	var secretKey []byte
	switch userType {
	case User:
		secretKey = []byte(cfg.Jwt.SecretKey)
	case Admin:
		secretKey = []byte(cfg.Jwt.AdminSecretKey)
	default:
		return nil, wrapper.InternalServerError("No matching role!")
	}

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
