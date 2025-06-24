package storage

import (
	"webapp/models"

	"gorm.io/gorm"
)

type UserStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) UserStorage {
	return UserStorage{
		db: db,
	}
}

func (s UserStorage) All() ([]models.UserDBModel, error) {
	var users []models.UserDBModel

	result := s.db.Preload("Emails").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
