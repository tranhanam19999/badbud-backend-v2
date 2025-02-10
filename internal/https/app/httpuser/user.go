package bdhttpuser

import "github.com/labstack/echo/v4"

func RegisterHttpUser(svc UserService, e *echo.Group) {
	h := HttpUser{svc: svc}

	e.GET("", h.list)
}

func (h *HttpUser) list(c echo.Context) error {
	return nil
}
