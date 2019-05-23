package services

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
)

type AuthService struct {
	authRepo repositories.IAuthRepo
}

func NewAuthService(authRepo repositories.IAuthRepo) *AuthService {
	return &AuthService{authRepo}
}

func (s *AuthService) Signin(email, password string) (*entities.User, error) {
	return s.authRepo.Signin(email, password)
}