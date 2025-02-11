package service

import (
	"errors"

	"github.com/badbud-backend-v2/internal/model"
	"github.com/badbud-backend-v2/internal/repo"
	"github.com/badbud-backend-v2/internal/service/dto"
	"gorm.io/gorm"
)

type Auth struct {
	Base
	repos      *repo.Repository
	jwtService JWT
}

func NewAuth(repos *repo.Repository, jwtService JWT) *Auth {
	return &Auth{
		jwtService: jwtService,
		repos:      repos,
	}
}

func (a *Auth) Login(ctx BDContext, input *dto.LoginReq) (*dto.LoginResp, error) {
	if err := a.Validate(ctx, input); err != nil {
		return nil, err
	}

	usr, err := a.repos.User.FindByUsername(input.Username)
	if err != nil {
		return nil, err
	}

	token, err := a.jwtService.GenerateToken(usr.ID, usr.Username)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResp{
		Token: token,
	}, nil
}

func (a *Auth) Register(ctx BDContext, input *dto.RegisterReq) (*dto.RegisterResp, error) {
	if err := a.Validate(ctx, input); err != nil {
		return nil, err
	}

	usr, err := a.repos.User.FindByUsername(input.Username)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		// Create user
		if err := a.repos.User.Create(&model.User{
			Username: input.Username,
			Name:     input.Name,
		}); err != nil {
			return nil, err
		}

		freshUser, err := a.repos.User.FindByUsername(input.Username)
		if err != nil {
			return nil, err
		}

		usr.ID = freshUser.ID
		usr.Username = freshUser.Username
	}

	token, err := a.jwtService.GenerateToken(usr.ID, usr.Username)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterResp{
		Token: token,
	}, nil
}
