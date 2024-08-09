package model

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	Name  string
	Price uint
}
