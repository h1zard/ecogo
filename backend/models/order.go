package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CartID uint
	UserID uint
}
