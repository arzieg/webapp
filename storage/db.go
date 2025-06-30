package storage

import (
	"gorm.io/gorm"
)

type UserDBModel struct {
	gorm.Model                  // add Columns ID, CreatedAt,UpdatedAt
	FirstName    string         `gorm:"column:first_name"`
	LastName     string         `gorm:"column:last_name"`
	Emails       []EmailDBModel `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	PasswordHash string         `gorm:"column:password_hash"`
	LastIP       string         `gorm:"column:last_ip"`
}

func (UserDBModel) TableName() string {
	return "users"
}

type EmailDBModel struct {
	gorm.Model
	Address string `gorm:"column:address;size:256;uniqueIndex"`
	Primary bool   `gorm:"column:primary"`
	UserID  int    `gorm:"column:user_id"`
}

type CreateUserRequestDBModel struct {
	FirstName string `json:"first_name" validate:"required_without=LastName"`
	LastName  string `json:"last_name" validate:"required_without=FirstName"`
	Email     string `json:"email" validate:"required,email"`
}

func (EmailDBModel) TableName() string {
	return "emails"
}

func AutoMigrate(db *gorm.DB) error {

	err := db.AutoMigrate(&UserDBModel{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&EmailDBModel{})
	if err != nil {
		return err
	}
	return nil
}
