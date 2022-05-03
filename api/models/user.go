package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		ID uint
		Username string
		Email *string
		FirstName string
		LastName string
		Password string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)