package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/Snake1-1eyes/todo-app"
	"github.com/Snake1-1eyes/todo-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const TokenTTL = 12 * time.Hour

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthServise struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthServise {
	return &AuthServise{repo: repo}
}

func (s *AuthServise) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthServise) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
