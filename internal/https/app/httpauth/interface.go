package bdhttpauth

import (
	bdhttp "github.com/badbud-backend-v2/internal/https"
	"github.com/badbud-backend-v2/internal/service"
	"github.com/badbud-backend-v2/internal/service/dto"
)

type HttpAuth struct {
	bdhttp.Base
	svc AuthService
}

type AuthService interface {
	Login(ctx service.BDContext, input *dto.LoginReq) (*dto.LoginResp, error)
	Register(ctx service.BDContext, input *dto.RegisterReq) (*dto.RegisterResp, error)
}
