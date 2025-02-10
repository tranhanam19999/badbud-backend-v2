package service

import (
	"errors"

	"github.com/badbud-backend-v2/internal/model"
	"github.com/badbud-backend-v2/internal/repo"
	"github.com/badbud-backend-v2/internal/service/dto"
)

type Match struct {
	Base
	repos *repo.Repository
}

func NewMatch(repos *repo.Repository) *Match {
	return &Match{
		repos: repos,
	}
}

func (m *Match) List(ctx BDContext, req *dto.ListMatchReq) (*dto.ListMatchResp, error) {
	if err := m.Validate(ctx, req); err != nil {
		return nil, err
	}

	total, matches, err := m.repos.Match.List(req.Page, req.Limit)
	if err != nil {
		return nil, err
	}

	return &dto.ListMatchResp{
		Total: total,
		Items: matches,
	}, nil
}

// Create match for people request to play
func (m *Match) Create(ctx BDContext, req *dto.CreateMatchReq) error {
	if err := m.Validate(ctx, req); err != nil {
		return err
	}

	// Check if court is available
	_, err := m.repos.Court.FindByID(req.CourtID)
	if err != nil {
		return err
	}

	// Validate time, should move to struct validate
	if req.StartTime.After(req.EndTime) {
		return errors.New("Start time can't be less then end time")
	}

	err = m.repos.Match.Create(&model.Match{
		CourtID:   req.CourtID,
		FeeMale:   req.Fee.Male,
		FeeFemale: req.Fee.Female,
	})

	return err
}

func (m *Match) CreateMatchRequest(ctx BDContext, req *dto.CreateMatchRequestReq) error {
	return nil
}

func (m *Match) AcceptMatchRequest(ctx BDContext, req *dto.AcceptMatchRequestReq) error {
	return nil

}

func (m *Match) RejectMatchRequest(ctx BDContext, req *dto.RejectMatchRequestReq) error {
	return nil
}
