package models

import (
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	OwnerID  uint
	TargetID uint
	Type     int // 1好友 2群組 3
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}

func SearchFriend(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id=? and type = 1", userId).Find(&contacts)
	// fmt.Println("contacts", userId)
	for _, v := range contacts {
		// fmt.Println("vvvvvv", v)
		objIds = append(objIds, uint64(v.TargetID))
	}
	fmt.Println("objIds", objIds)
	users := make([]UserBasic, 0)
	utils.DB.Where("id in ?", objIds).Find(&users)
	return users
}
