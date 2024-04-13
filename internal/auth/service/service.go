package service

import (
	"context"

	"github.com/RickDred/ozinse/internal/auth"
	"github.com/RickDred/ozinse/internal/models"
)

type service struct {
	repo auth.RepoInterface
}

func New(repo auth.RepoInterface) auth.ServiceInterface {
	return &service{repo}
}

func (s *service) Register(ctx context.Context, user *models.User) (string, error) {
	user.Standardize()

	if err := user.Validate(user.ValidateEmail, user.ValidatePassword); err != nil {
		return "", err
	}

	if err := user.HashPassword(); err != nil {
		return "", err
	}

	_, err := s.repo.Insert(ctx, user)
	if err != nil {
		return "", err
	}

	user.CleanPassword()

	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) Login(ctx context.Context, input *models.User) (string, error) {
	user, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	if err := user.ComparePassword(input.Password); err != nil {
		return "", err
	}

	user.CleanPassword()

	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
