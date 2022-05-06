package models

import (
	"time"
)

type (
	User struct {
		ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
		Username  string    `gorm:"type:string;unique;not null" json:"username"`
		Email     string    `gorm:"type:string;unique;not null" json:"email"`
		FirstName string    `gorm:"type:string;not null" json:"first_name"`
		LastName  string    `gorm:"type:string;not null" json:"last_name"`
		OtherName string    `gorm:"type:string" json:"other_name"`
		Password  string    `gorm:"type:string" json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `gorm:"index" json:"deleted_at"`
		Contacts  []Contact `json:"contacts,omitempty"`
	}
)
