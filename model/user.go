package model

import (
	"ecommerce/database"
	"html"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username" validate:"required,min=2,max=30"`
	// Email    string  `gorm:"size:255;not null;"       json:"email"    validate:"email,required"`
	Password string `gorm:"size:255;not null;"       json:"-"        validate:"required,min=6"`
	// Phone    *string `gorm:"size:255;not null;"       json:"phone"    validate:"required"`
	Address []Address
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}
func (user *User) ValidatePassword(password string) error {
	log.Println(password)
	log.Println(user.Password)
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Address").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
