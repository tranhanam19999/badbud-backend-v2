package repo

import "gorm.io/gorm"

type Repository struct {
	db           *gorm.DB
	User         *UserRepo
	Match        *MatchRepo
	MatchRequest *MatchRequestRepo
	Court        *CourtRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:           db,
		User:         NewUserRepo(db),
		Match:        NewMatchRepo(db),
		MatchRequest: NewMatchRequestRepo(db),
		Court:        NewCourtRepo(db),
	}
}

// A wrapper transaction functions so that services which inject repository can use
func (r *Repository) Transaction(txFunc func(txRepo *Repository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		txRepo := &Repository{
			db:           tx,
			User:         &UserRepo{db: tx},
			Match:        &MatchRepo{db: tx},
			MatchRequest: &MatchRequestRepo{db: tx},
			Court:        &CourtRepo{db: tx},
		}
		return txFunc(txRepo)
	})
}
