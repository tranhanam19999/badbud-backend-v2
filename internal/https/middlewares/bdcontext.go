package middlewares

import (
	bdhttp "github.com/badbud-backend-v2/internal/https"
	"github.com/labstack/echo/v4"
)

func BDContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bdCtx := &bdhttp.BDContext{Context: c} // Wrap Echo’s Context
		return next(bdCtx)                     // Pass the new context
	}
}
