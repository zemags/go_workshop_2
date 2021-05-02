package service

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/zemags/go_workshop_2/pkg/repository"
	workshop_2 "github.com/zemags/go_workshop_2/store"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	// to work with db need repo
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user workshop_2.User) (int, error) {
	// pass User struct to lower level, to repository
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
