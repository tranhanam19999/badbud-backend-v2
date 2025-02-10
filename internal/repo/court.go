package repo

import (
	"github.com/badbud-backend-v2/internal/model"
	"gorm.io/gorm"
)

type CourtRepo struct {
	db *gorm.DB
}

func NewCourtRepo(db *gorm.DB) *CourtRepo {
	return &CourtRepo{db: db}
}

func (m *CourtRepo) FindByID(id int) (court *model.Court, err error) {
	err = m.db.Find(&court, id).Error
	return
}
