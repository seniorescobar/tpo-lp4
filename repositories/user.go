package repositories

import "github.com/stretchr/testify/mock"

type IUserRepo interface {
	Register(email, password string) error
}

type UserRepo struct{}

func (r *UserRepo) Register(email, password string) error {
	// register user
	// ...
	return nil
}

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) Register(email, password string) error {
	args := m.Called(email, password)
	return args.Error(0)
}
