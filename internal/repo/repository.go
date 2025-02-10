package repo

import "gorm.io/gorm"

type Repository struct {
	User         *UserRepo
	Match        *MatchRepo
	MatchRequest *MatchRequestRepo
	Court        *CourtRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:         NewUserRepo(db),
		Match:        NewMatchRepo(db),
		MatchRequest: NewMatchRequestRepo(db),
		Court:        NewCourtRepo(db),
	}
}
