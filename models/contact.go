package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerID  uint
	TargetID uint
	Type     int
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
