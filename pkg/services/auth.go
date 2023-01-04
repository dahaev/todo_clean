package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dahaev/todo.git"
	"github.com/dahaev/todo.git/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt     = "093021dksak21j[]sa"
	tokenTTL = 12 * time.Hour
	signKey  = "HJDSA9dsaklaHS"
)

type tokenClaims struct {
	jwt.StandardClaims
	Userid int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}

}

func (auth *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return auth.repo.CreateUser(user)
}

func (auth *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims not of type")
	}

	return claims.Userid, nil
}

func (auth *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := auth.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signKey))

}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
