package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Snake1-1eyes/todo-app"
	"github.com/Snake1-1eyes/todo-app/pkg/repository"
)

const salt = "sldcm23ke2mssdc"

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

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
