package storage

import (
	"webapp/domain"

	"gorm.io/gorm"
)

type UserStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) All() (*[]domain.User, error) {
	var users []UserDBModel

	result := s.db.Preload("Emails").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	domainUsers := make([]domain.User, len(users))
	for i, u := range users {
		domainUsers[i] = toUserDomainModel(u)
	}

	return &domainUsers, nil
}
