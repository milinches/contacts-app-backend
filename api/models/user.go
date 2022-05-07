package models

import (
	"time"

	"github.com/milinches/contacts-app-backend/api/utils"
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
		Contacts  []Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"contacts"`
	}
)

func BeforeSave(u *User) error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
