package models

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Username  string    `gorm:"type:varchar(20);unique;not null" json:"username"`
		Email     string    `gorm:"type:varchar(40);unique;not null" json:"email"`
		FirstName string    `gorm:"type:varchar(40);not null" json:"first_name"`
		LastName  string    `gorm:"type:varchar(40);not null" json:"last_name"`
		OtherName string    `gorm:"type:varchar(40)" json:"other_name"`
		Password  string    `gorm:"type:varchar(30)" json:"password"`
		Contacts  []Contact `json:"contacts,omitempty"`
	}
)
