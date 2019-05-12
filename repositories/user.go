package repositories

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type IUserRepo interface {
	Register(email, password string) error
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Register(email, password string) error {
	// register user
	// ...
	return nil
}

type UserRepoMock struct {
	mock.Mock
}

func NewUserRepoMock() *UserRepoMock {
	return new(UserRepoMock)
}

func (m *UserRepoMock) Register(email, password string) error {
	args := m.Called(email, password)
	return args.Error(0)
}
