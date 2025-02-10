package repo

import "gorm.io/gorm"

// Match Repository
type MatchRequestRepo struct {
	db *gorm.DB
}

func NewMatchRequestRepo(db *gorm.DB) *MatchRequestRepo {
	return &MatchRequestRepo{db: db}
}

// func (r *MatchRepo) CreateMatch(match *Match) error {
// 	return r.db.Create(match).Error
// }

// func (r *MatchRepo) GetMatch(id uint) (*Match, error) {
// 	var match Match
// 	result := r.db.Preload("User").First(&match, id)
// 	return &match, result.Error
// }
