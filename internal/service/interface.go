package service

import (
	"context"

	"github.com/badbud-backend-v2/internal/service/dto"
)

type BDContext interface {
	AuthUser() *dto.AuthUser
	ReqCtx() context.Context
}
