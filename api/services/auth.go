package services

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
)

type AuthService struct {
	authRepo repositories.IUserRepo
}

func NewAuthService(authRepo repositories.IUserRepo) *AuthService {
	return &AuthService{authRepo}
}

func (s *AuthService) Register(u *entities.User) error {
	return s.authRepo.Register(u)
}

func (s *AuthService) Signin(email, password string) (*entities.User, error) {
	return s.authRepo.Signin(email, password)
}
