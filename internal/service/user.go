package service

import (
	"github.com/badbud-backend-v2/internal/repo"
	"github.com/badbud-backend-v2/internal/service/dto"
)

type User struct {
	repos *repo.Repository
}

func NewUser(repos *repo.Repository) *User {
	return &User{
		repos: repos,
	}
}

func (m *User) List(ctx BDContext, req *dto.ListUserReq) (*dto.ListUserResp, error) {
	return nil, nil
}
