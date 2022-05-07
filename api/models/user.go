package models

import (
	"time"
)

type (
	User struct {
		ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
		Email     string    `gorm:"type:string;uniqueIndex;not null" json:"email"`
		FirstName string    `gorm:"type:string;not null" json:"first_name"`
		LastName  string    `gorm:"type:string;not null" json:"last_name"`
		Password  string    `gorm:"type:string;not null" json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `gorm:"index" json:"deleted_at"`
		Contact   []Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"contact,omitempty"`
	}
)

func (u *User) TableName() string {
	return "user"
}
