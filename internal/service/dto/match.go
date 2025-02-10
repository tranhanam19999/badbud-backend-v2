package dto

import (
	"time"

	"github.com/badbud-backend-v2/internal/model"
)

type ListMatchReq struct {
	Page  int
	Limit int
}

type ListMatchResp struct {
	Total int64
	Items []*model.Match
}

type CreateMatchReq struct {
	CourtID   int
	Fee       FeeReq
	StartTime time.Time
	EndTime   time.Time
}

type FeeReq struct {
	Male   string
	Female string
}
