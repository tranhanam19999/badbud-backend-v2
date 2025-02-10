package service

import (
	"context"

	"github.com/badbud-backend-v2/internal/common/errors"
	"github.com/badbud-backend-v2/internal/common/validate"
	"github.com/go-playground/validator/v10"
)

type customValidate interface {
	Validate(c context.Context) *errors.Error
}

type Base struct{}

func (b *Base) Validate(ctx BDContext, data any) error {
	err := validate.Struct(data)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		return errs
	}

	if v, ok := data.(customValidate); ok {
		return v.Validate(ctx.ReqCtx())
	}

	return nil
}
