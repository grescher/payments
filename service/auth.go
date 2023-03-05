package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"payments/config"
	"payments/models"
	"payments/repository"

	"github.com/zeebo/errs"
)

var authErr = errs.Class("authorization service")

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(r *repository.Repository) *AuthorizationService {
	return &AuthorizationService{repo: r}
}

func (s *AuthorizationService) CreateUser(ctx context.Context, user models.User) (id int, err error) {
	user.PasswordHash = generatePasswordHash(user.PasswordHash)
	if id, err = s.repo.CreateUser(ctx, user); err != nil {
		return 0, authErr.Wrap(err)
	}
	log.Printf("Service: You've created user with id %v\n", id)

	return id, nil
}

func generatePasswordHash(password string) (passwordHash string) {
	hash := sha256.New()
	hash.Write([]byte(password))
	passwordHash = fmt.Sprintf("%x", hash.Sum([]byte(config.HashSalt())))
	return passwordHash
}
