package services

import (
	"bitbucket.org/aj5110/tpo-lp4/repositories"
)

type AuthService struct {
	userRepo repositories.IUserRepo
}

func NewAuthService(userRepo repositories.IUserRepo) *AuthService {
	return &AuthService{userRepo}
}

func (s *AuthService) RegisterUser(email, password string) error {
	return s.userRepo.Register(email, password)
}
