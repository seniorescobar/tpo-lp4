package services

import (
	"bitbucket.org/aj5110/tpo-lp4/api/entities"
	"bitbucket.org/aj5110/tpo-lp4/api/repositories"
)

type AuthService struct {
	userRepo repositories.IUserRepo
}

func NewAuthService(userRepo repositories.IUserRepo) *AuthService {
	return &AuthService{userRepo}
}

func (s *AuthService) Register(u *entities.User) (*entities.User, error) {
	return s.userRepo.Register(u)
}

func (s *AuthService) Signin(email, password string) (*entities.User, error) {
	return s.userRepo.Signin(email, password)
}

func (s *AuthService) ChangePassword(email, oldPassword, newPassword string) error {
	return s.userRepo.ChangePassword(email, oldPassword, newPassword)
}
