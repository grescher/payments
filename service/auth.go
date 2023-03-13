package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"payments/config"
	"payments/models"
	"payments/repository"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/zeebo/errs"
)

var authErr = errs.Class("authorization service")

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID int `json:"user_id"`
}

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
	return id, nil
}

func (s *AuthorizationService) GenerateToken(ctx context.Context, user models.User) (signedToken string, err error) {
	id, err := s.repo.GetUserID(ctx, user.Email, generatePasswordHash(user.PasswordHash))
	if err != nil {
		return "", authErr.Wrap(err)
	}

	claims := tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.TokenTTL())),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if signedToken, err = token.SignedString([]byte(config.TokenSignature())); err != nil {
		return "", authErr.Wrap(err)
	}

	return signedToken, nil
}

func generatePasswordHash(password string) (passwordHash string) {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(config.HashSalt())))
}
