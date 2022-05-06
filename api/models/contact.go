package models

import "gorm.io/gorm"

type (
	Contact struct {
		gorm.Model
		Email       string   `gorm:"type:varchar(40)" json:"email"`
		FirstName   string   `gorm:"type:varchar(40);not null" json:"first_name"`
		LastName    string   `gorm:"type:varchar(40)" json:"last_name"`
		OtherName   string   `gorm:"type:varchar(40)" json:"other_name"`
		Address     string   `gorm:"type:text;" json:"address"`
		PhoneNumber []string `gorm:"type:varchar(20)" json:"phone_number"`
		User        User     `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"user"`
		UserID      uint     `gorm:"not null" json:"user_id"`
	}
)
