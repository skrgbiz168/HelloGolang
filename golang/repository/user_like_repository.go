package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IUserLikeRepository interface {
	GetUsers(user *model.User, id uint) error
}

type userLikeRepository struct {
	db *gorm.DB
}

func NewUserLikeRepository(db *gorm.DB) IUserLikeRepository {
	return &userLikeRepository{db}
}
// GetUsers is a function to get all users
func (ur *userLikeRepository) GetUsers(user *model.User, id uint) error {
	// User以外のユーザーを取得
	if err := ulr.db.Where("id!=?", id).Find(user).Error; err != nil {
		return err
	}
	return nil
}

