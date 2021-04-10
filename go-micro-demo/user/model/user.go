package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string  `json:"name" gorm:"type:char(13)"`
	Phone string `json:"phone" gorm:"type:char(11)"`
	Password string  `json:"password" gorm:"type:char(13)"`
}


type Loginer struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type UpdatePassword struct {
	Uid uint32 `json:"uid"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}