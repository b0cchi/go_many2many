package model

import "gorm.io/gorm"

type Dept struct {
	gorm.Model
	Name string

	Users []*User `gorm:"many2many:user_depts"`
}
