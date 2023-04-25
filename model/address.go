package model

import (
	"ecommerce/database"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Name       string `json:"address_name"`
	Street     string `json:"street_name"`
	Complement string `json:"complement"`
	City       string `json:"city_name"`
	CEP        string `json:"cep"`
	UserID     uint
}

func (address *Address) Save() (*Address, error) {
	err := database.Database.Create(&address).Error
	if err != nil {
		return &Address{}, err
	}
	return address, nil
}
