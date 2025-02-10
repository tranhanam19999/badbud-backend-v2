package bdhttpmatch

import (
	bdhttp "github.com/badbud-backend-v2/internal/https"
	"github.com/badbud-backend-v2/internal/service"
	"github.com/badbud-backend-v2/internal/service/dto"
)

type HttpMatch struct {
	bdhttp.Base
	svc MatchService
}

type MatchService interface {
	List(ctx service.BDContext, req *dto.ListMatchReq) (*dto.ListMatchResp, error)
	Create(ctx service.BDContext, req *dto.CreateMatchReq) error

	CreateMatchRequest(ctx service.BDContext, req *dto.CreateMatchRequestReq) error
	AcceptMatchRequest(ctx service.BDContext, req *dto.AcceptMatchRequestReq) error
	RejectMatchRequest(ctx service.BDContext, req *dto.RejectMatchRequestReq) error
}
