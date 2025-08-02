package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID uint
	User   User
	Items  []CartItem // NOT Item
	Status string
}
