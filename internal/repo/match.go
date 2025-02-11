package repo

import (
	"github.com/badbud-backend-v2/internal/model"
	"gorm.io/gorm"
)

// Match Repository
type MatchRepo struct {
	db *gorm.DB
}

func NewMatchRepo(db *gorm.DB) *MatchRepo {
	return &MatchRepo{db: db}
}

func (r *MatchRepo) Create(match *model.Match) error {
	return r.db.Create(match).Error
}

func (r *MatchRepo) FindByID(id string) (match *model.Match, err error) {
	err = r.db.Preload("MatchParticipants").Find(&match, id).Error
	return
}

func (r *MatchRepo) List(page, limit int) (total int64, matches []*model.Match, err error) {
	err = r.db.Offset(page - 1).Limit(limit).Find(&matches).Error
	if err != nil {
		return 0, nil, err
	}

	err = r.db.Offset(page - 1).Limit(limit).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	return
}

func (r *MatchRepo) Update(values any, conds ...any) error {
	err := r.db.Where(conds).Updates(values).Error
	return err
}

// AppendParticipant adds a participant to a match
func (r *MatchRepo) AddParticipant(matchId, userId string) error {
	participant := &model.MatchParticipant{
		MatchId: matchId,
		UserId:  userId,
	}

	// Check if the participant already exists
	var existing model.MatchParticipant
	err := r.db.Where("match_id = ? AND user_id = ?", matchId, userId).First(&existing).Error
	if err == nil {
		return nil // Already exists, no need to insert
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	// Append participant
	return r.db.Create(&participant).Error
}

// RemoveParticipant removes a participant from a match
func (r *MatchRepo) RemoveParticipant(matchId, userId string) error {
	return r.db.Where("match_id = ? AND user_id = ?", matchId, userId).Delete(&model.MatchParticipant{}).Error
}
