package bdhttp

import "github.com/labstack/echo/v4"

type Binder = echo.DefaultBinder

type Base struct {
	binder Binder
}

func (b *Base) Context(e echo.Context) *BDContext {
	return e.(*BDContext)
}

func (b *Base) Binder() *Binder {
	return &b.binder
}
