package storage

import (
	"webapp/domain"

	"gorm.io/gorm"
)

func toEmailDBModel(e domain.Email) EmailDBModel {
	return EmailDBModel{
		Model:   gorm.Model{ID: e.ID},
		Address: e.Address,
		Primary: e.Primary,
		UserID:  e.UserID,
	}
}

func toEmailDomainModel(m EmailDBModel) domain.Email {
	return domain.Email{
		ID:      m.ID,
		Address: m.Address,
		Primary: m.Primary,
		UserID:  m.UserID,
	}
}

// helper function to create a slice
func toEmailDBModels(emails []domain.Email) []EmailDBModel {
	dbEmails := make([]EmailDBModel, len(emails))
	for i, e := range emails {
		dbEmails[i] = toEmailDBModel(e)
	}
	return dbEmails
}

func toUserDomainModel(m UserDBModel) domain.User {
	emails := make([]domain.Email, len(m.Emails))
	for i, e := range m.Emails {
		emails[i] = domain.Email{
			ID:      e.ID,
			Address: e.Address,
			Primary: e.Primary,
			UserID:  e.UserID,
		}
	}

	return domain.User{
		ID:        m.ID,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Emails:    emails,
		LastIP:    m.LastIP,
	}
}

func toUserDBModel(u domain.User) UserDBModel {
	return UserDBModel{
		Model:     gorm.Model{ID: u.ID},
		FirstName: u.FirstName,
		LastName:  u.LastName,
		LastIP:    u.LastIP,
		Emails:    toEmailDBModels(u.Emails),
	}
}

// func toCreateUserRequestDBModel(u domain.CreateUserRequest) CreateUserRequestDBModel {
// 	return CreateUserRequestDBModel{
// 		FirstName: u.FirstName,
// 		LastName:  u.LastName,
// 		Email:     u.Email,
// 	}
// }

// func toCreateUserRequestDomainModel(u CreateUserRequestDBModel) domain.CreateUserRequest {
// 	return domain.CreateUserRequest{
// 		FirstName: u.FirstName,
// 		LastName:  u.LastName,
// 		Email:     u.Email,
// 	}
// }
