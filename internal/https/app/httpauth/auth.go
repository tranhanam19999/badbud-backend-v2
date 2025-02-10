package bdhttpauth

import (
	"net/http"

	"github.com/badbud-backend-v2/internal/service/dto"
	"github.com/labstack/echo/v4"
)

func RegisterHttpAuth(svc AuthService, e *echo.Group) {
	h := HttpAuth{svc: svc}

	e.GET("", h.login)
	e.POST("", h.register)

}

func (h *HttpAuth) login(c echo.Context) error {
	input := &dto.LoginReq{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	data, err := h.svc.Login(h.Context(c), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *HttpAuth) register(c echo.Context) error {
	input := &dto.RegisterReq{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	data, err := h.svc.Register(h.Context(c), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
