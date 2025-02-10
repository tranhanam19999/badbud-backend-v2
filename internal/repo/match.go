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

func (r *MatchRepo) FindByID(id uint) (match *model.Match, err error) {
	err = r.db.Find(&match, id).Error
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
