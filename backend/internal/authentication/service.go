package authentication

import (
	"box/internal/user"
	"context"
	"errors"
)

var ErrFailedLoginAttempt = errors.New("Failed login attempt")

type Service interface {
	Login(ctx context.Context, username, password string) (token *Token, user *user.User, err error)
}

type ServiceImpl struct{}

func New() Service {
	return &ServiceImpl{}
}

func (s *ServiceImpl) Login(ctx context.Context, username, password string) (*Token, *user.User, error) {
	return &Token{}, &user.User{}, nil
}
