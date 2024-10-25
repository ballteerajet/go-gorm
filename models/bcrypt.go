package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
