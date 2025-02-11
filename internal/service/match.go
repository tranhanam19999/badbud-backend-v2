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
		CourtNum:  req.CourtNum,
		FeeMale:   req.Fee.Male,
		FeeFemale: req.Fee.Female,
		Limit:     req.Limit,
	})

	return err
}

func (m *Match) CreateMatchRequest(ctx BDContext, req *dto.CreateMatchRequestReq) error {
	curMatch, err := m.repos.Match.FindByID(req.MatchId)
	if err != nil {
		return err
	}

	if len(curMatch.MatchParticipants)+1 > curMatch.Limit {
		return errors.New("Participants limit exceeded")
	}

	// TODO: Check if requested user is banned, not verify email, otp,...

	if err = m.repos.Transaction(func(tx *repo.Repository) error {
		if err := tx.MatchRequest.Create(&model.MatchRequest{
			UserID:  ctx.AuthUser().ID,
			MatchID: req.MatchId,
			Status:  model.MatchRequestStatusRequested,
		}); err != nil {
			return err
		}

		// TODO: Notification?
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// Is for owner to accept the match request send by the player
func (m *Match) AcceptMatchRequest(ctx BDContext, req *dto.AcceptMatchRequestReq) error {
	matchRequest, err := m.repos.MatchRequest.FindByID(req.RequestId)
	if err != nil {
		return err
	}

	if matchRequest.Status != model.MatchRequestStatusRequested {
		return errors.New("Invalid match request status")
	}

	match, err := m.repos.Match.FindByID(matchRequest.MatchID)
	if err != nil {
		return err
	}

	if err := m.repos.Transaction(func(txRepo *repo.Repository) error {
		if err := txRepo.MatchRequest.Update(map[string]any{
			"status": model.MatchRequestStatusAccepted,
		}, "id = ?", req.RequestId); err != nil {
			return err
		}

		if err := txRepo.Match.AddParticipant(matchRequest.MatchID, matchRequest.UserID); err != nil {
			return err
		}

		// If after accept the limit execeeded -> Auto reject all other participant
		// Plus 1 since we accepted the match request
		if len(match.MatchParticipants)+1 == match.Limit {
			// Auto reject all others match request
			if err := txRepo.MatchRequest.Update(map[string]any{
				"status": model.MatchRequestStatusRejected,
			}, "court_id = ? AND court_num = ?", match.CourtID, match.CourtNum); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (m *Match) RejectMatchRequest(ctx BDContext, req *dto.RejectMatchRequestReq) error {
	matchRequest, err := m.repos.MatchRequest.FindByID(req.RequestId)
	if err != nil {
		return err
	}

	if matchRequest.Status != model.MatchRequestStatusRequested {
		return errors.New("Invalid match request status")
	}

	if err := m.repos.Transaction(func(txRepo *repo.Repository) error {
		if err := txRepo.MatchRequest.Update(map[string]any{
			"status": model.MatchRequestStatusAccepted,
		}, "id = ?", req.RequestId); err != nil {
			return err
		}

		if err := txRepo.Match.RemoveParticipant(matchRequest.MatchID, matchRequest.UserID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
