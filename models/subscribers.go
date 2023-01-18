package models

import (
	"gorm.io/gorm"
)

type Subscribers struct {
	gorm.Model
	Name  string
	Email string
}
