package service

import (
	"crypto/sha1"
	"fmt"
	// "hash"

	todo "github.com/KDias-code/todoapp"
	"github.com/KDias-code/todoapp/pkg/repository"
)

const salt = "asdardirqweo231koads"

type AuthService struct {
	repo repository.Autharization
}

func NewAuthService(repo repository.Autharization) *AuthService{
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error){
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string{
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x" , hash.Sum([]byte(salt)))
}