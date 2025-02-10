package repo

import (
	"github.com/badbud-backend-v2/internal/model"
	"gorm.io/gorm"
)

// Match Repository
type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) FindByID(id int) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) Create(input *model.User) error {
	if err := r.db.Create(&input).Error; err != nil {
		return err
	}

	return nil
}
