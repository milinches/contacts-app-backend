package models

import (
	"time"

	"github.com/lib/pq"
)

type (
	Contact struct {
		ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
		Email       string         `gorm:"type:string" json:"email"`
		Name        string         `gorm:"type:string;not null" json:"name"`
		Address     string         `gorm:"type:text;" json:"address"`
		PhoneNumber pq.StringArray `gorm:"type:text[];not null" json:"phone_number"`
		Type        pq.StringArray `gorm:"type:text[];not null" json:"type"`
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
		DeletedAt   time.Time      `gorm:"index" json:"deleted_at"`
		UserID      uint           `json:"user_id"`
	}
)

// Added Hooks — Gorm hooks are function that are called before or after creation

func (u *Contact) TableName() string {
	return "contact"
}
