package service

import (
	"context"

	"github.com/awleory/kode/notebook/internal/entity"
)

type Hash interface {
	Password(pass string) (string, error)
}

type UserRepo interface {
	CreateUser(ctx context.Context, user entity.SignUpInput) error
	GetUser(ctx context.Context, email, pass string) (int, error)
}

type UserService struct {
	repo UserRepo
	hash Hash
}

func NewUsers(repo UserRepo, passHash Hash) *UserService {
	return &UserService{
		repo: repo,
		hash: passHash,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user entity.SignUpInput) error {
	passHash, err := s.hash.Password(user.Password)
	if err != nil {
		return err
	}

	user.Password = passHash
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) VerifyUser(ctx context.Context, email, pass string) (int, error) {
	passHash, err := s.hash.Password(pass)
	if err != nil {
		return -1, err
	}

	userId, err := s.repo.GetUser(ctx, email, passHash)
	if err != nil {
		return -1, err
	}

	return userId, err
}
