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
		Contact   []*Contact `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"contact,omitempty"`
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

func (u *User) Prepare() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate() (err error) {
	if u.Email == "" {
		return errors.New("Required email")
	}
	if u.FirstName == "" {
		return errors.New("Required firstname")
	}
	if u.LastName == "" {
		return errors.New("Required lastname")
	}
	if u.Password == "" {
		return errors.New("Required password")
	}
	return
}
