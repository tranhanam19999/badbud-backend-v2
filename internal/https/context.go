package bdhttp

import (
	"context"

	"github.com/badbud-backend-v2/internal/service"
	"github.com/badbud-backend-v2/internal/service/dto"
	"github.com/labstack/echo/v4"
)

type BDContext struct {
	echo.Context
	user *dto.AuthUser
}

var _ service.BDContext = &BDContext{}

func (bd *BDContext) AuthUser() *dto.AuthUser {
	return bd.user
}

func (bd *BDContext) ReqCtx() context.Context {
	return bd.Request().Context()
}

func (bd *BDContext) SetAuthUser(u *dto.AuthUser) {
	bd.user = u
}
