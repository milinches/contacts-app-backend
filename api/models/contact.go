package models

import (
	"time"

	"github.com/lib/pq"
)

type (
	Contact struct {
		ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
		Email       string         `gorm:"type:string" json:"email"`
		FirstName   string         `gorm:"type:string;not null" json:"first_name"`
		LastName    string         `gorm:"type:string" json:"last_name"`
		OtherName   string         `gorm:"type:string" json:"other_name"`
		Address     string         `gorm:"type:text;" json:"address"`
		PhoneNumber pq.StringArray `gorm:"type:text[]" json:"phone_number"`
		User        User           `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"user"`
		UserID      uint           `gorm:"not null" json:"userID"`
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
		DeletedAt   time.Time      `gorm:"index" json:"deleted_at"`
	}
)
