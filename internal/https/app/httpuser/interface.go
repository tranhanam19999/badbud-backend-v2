package bdhttpuser

import (
	bdhttp "github.com/badbud-backend-v2/internal/https"
	"github.com/badbud-backend-v2/internal/service"
	"github.com/badbud-backend-v2/internal/service/dto"
)

type HttpUser struct {
	bdhttp.Base
	svc UserService
}

type UserService interface {
	List(ctx service.BDContext, req *dto.ListUserReq) (*dto.ListUserResp, error)
}
