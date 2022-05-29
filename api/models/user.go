package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/milinches/contacts-app-backend/api/utils"
	"gorm.io/gorm"
)

type (
	User struct {
		ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
		Email     string     `gorm:"type:string;uniqueIndex;not null" json:"email"`
		FirstName string     `gorm:"type:string;not null" json:"first_name"`
		LastName  string     `gorm:"type:string;not null" json:"last_name"`
		Password  string     `gorm:"type:string;not null" json:"password"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt time.Time  `gorm:"index" json:"deleted_at"`
		Contact   []*Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"contact,omitempty"`
	}
)

// Added Hooks — Gorm hooks are function that are called before or after creation

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return errors.New("error hashing password")
	}
	u.Password = string(hashedPassword)
	return
}

func (user *User) Prepare() {
	user.ID = 0
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.FirstName = html.EscapeString(strings.TrimSpace(user.FirstName))
	user.LastName = html.EscapeString(strings.TrimSpace(user.LastName))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}
