package storage

import (
	"fmt"
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

func (s *UserStorage) Add(user *domain.User) error {

	return s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Omit("Emails").Create(&user)
		if result.Error != nil {
			return result.Error
		}

		email := &user.Emails[0]
		email.UserID = int(user.ID)

		result = tx.Create(&email)
		if result.Error != nil {
			fmt.Errorf("error creating user: %v", result.Error)
		}
		return nil
	})

}
