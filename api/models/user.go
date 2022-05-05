package models

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Username  string    `gorm:"size:20;uniqueIndex;not null" json:"username"`
		Email     *string   `gorm:"size:40;uniqueIndex;not null" json:"email"`
		FirstName string    `gorm:"size:40;not null" json:"first_name"`
		LastName  string    `gorm:"size:40;not null" json:"last_name"`
		OtherName string    `gorm:"size:40" json:"other_name"`
		Password  string    `gorm:"size:50" json:"password"`
		Contacts  []Contact `json:"contacts,omitempty"`
	}
)
