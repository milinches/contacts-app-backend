package models

import "gorm.io/gorm"

type (
	Contact struct {
		gorm.Model
		Email       *string  `gorm:"size:50" json:"email"`
		FirstName   string   `gorm:"size:50;not null" json:"first_name"`
		LastName    string   `gorm:"size:50" json:"last_name"`
		OtherName   string   `gorm:"size:50" json:"other_name"`
		Address     string   `gorm:"size:255;" json:"address"`
		PhoneNumber []string `gorm:"size:20" json:"phone_number"`
		User        User     `gorm:"foreignKey:UserID" json:"user"`
		UserID      uint     `gorm:"not null" json:"user_id"`
	}
)
