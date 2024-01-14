package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string

	Depts []*Dept `gorm:"many2many:user_depts"`
}
