package bdhttpmatch

import (
	"net/http"

	"github.com/badbud-backend-v2/internal/service/dto"
	"github.com/labstack/echo/v4"
)

func RegisterHttpMatch(svc MatchService, e *echo.Group) {
	h := HttpMatch{svc: svc}

	e.GET("", h.list)
	e.POST("", h.create)

	mrRouter := e.Group("/request")
	{
		mrRouter.GET("", h.listMatchRequest)
		mrRouter.POST("", h.createMatchRequest)
		mrRouter.POST("/accept", h.acceptMatchRequest)
		mrRouter.POST("/reject", h.rejectMatchRequest)
	}
}

func (h *HttpMatch) list(c echo.Context) error {
	input := &dto.ListMatchReq{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	datas, err := h.svc.List(h.Context(c), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, datas)
}

// Create a match for other users to join
func (h *HttpMatch) create(c echo.Context) error {
	input := &dto.CreateMatchReq{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.svc.Create(h.Context(c), input); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "Match created successfully")
}

func (h *HttpMatch) listMatchRequest(c echo.Context) error {
	return nil
}

// Create a match for other users to join
func (h *HttpMatch) createMatchRequest(c echo.Context) error {
	input := &dto.CreateMatchRequestReq{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.svc.CreateMatchRequest(h.Context(c), input); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "Match request created successfully")
}

func (h *HttpMatch) acceptMatchRequest(c echo.Context) error {
	input := &dto.AcceptMatchRequestReq{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.svc.AcceptMatchRequest(h.Context(c), input); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "Match request accepted successfully")
}

func (h *HttpMatch) rejectMatchRequest(c echo.Context) error {
	input := &dto.RejectMatchRequestReq{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.svc.RejectMatchRequest(h.Context(c), input); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "Match request rejected successfully")
}
