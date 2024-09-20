package users

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/logger"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/storage/users"
	"github.com/egor-zakharov/goph-keeper/internal/utils"
)

type service struct {
	storage users.Storage
}

func New(storage users.Storage) Service {
	return &service{storage: storage}
}

func (s *service) Register(ctx context.Context, userIn models.User) (*models.User, error) {
	hashPassword, err := utils.GetHashPassword(userIn.Password)
	if err != nil {
		logger.Log().Sugar().Errorw("Service register user", "get hash password", err)
		return nil, err
	}
	user := models.User{
		Login:    userIn.Login,
		Password: hashPassword,
	}
	registeredUser, err := s.storage.Register(ctx, user)
	if err != nil {
		logger.Log().Sugar().Errorw("Service register user", "user storage register", err)
		return nil, err
	}

	return registeredUser, nil
}

func (s *service) Login(ctx context.Context, userIn models.User) (*models.User, error) {
	user, err := s.storage.Login(ctx, userIn.Login)
	logger.Log().Sugar().Errorw("Service login user", "user storage login", err)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(user.Password, userIn.Password) {
		logger.Log().Sugar().Errorw("Service login user", "check password login", err)
		return nil, ErrIncorrectData
	}
	return user, nil
}
