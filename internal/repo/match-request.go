package repo

import (
	"github.com/badbud-backend-v2/internal/model"
	"gorm.io/gorm"
)

// Match Repository
type MatchRequestRepo struct {
	db *gorm.DB
}

func NewMatchRequestRepo(db *gorm.DB) *MatchRequestRepo {
	return &MatchRequestRepo{db: db}
}

func (r *MatchRequestRepo) Create(matchReq *model.MatchRequest) error {
	return r.db.Create(matchReq).Error
}

func (r *MatchRequestRepo) FindByID(id string) (matchRequest *model.MatchRequest, err error) {
	err = r.db.Find(&matchRequest, id).Error
	return
}

func (r *MatchRequestRepo) Update(values any, conds ...any) error {
	err := r.db.Where(conds).Updates(values).Error
	return err
}
