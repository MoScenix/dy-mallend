package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"type:varchar(255);not null;uniqueIndex"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
	UserName       string `gorm:"type:varchar(255);not null"`
	Description    string `gorm:"type:text"`
}

func (User) TableName() string {
	return "users"
}
func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
func GetByEmail(db *gorm.DB, email string) (*User, error) {
	var User User
	err := db.Where("email = ?", email).First(&User).Error
	return &User, err
}
