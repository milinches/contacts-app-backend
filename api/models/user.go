package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

func (u *User) Prepare() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			return errors.New("required email")
		}
		if u.FirstName == "" {
			return errors.New("required firstname")
		}
		if u.LastName == "" {
			return errors.New("required lastname")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	case "login":
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	default:
		if u.Email == "" {
			return errors.New("required email")
		}
		if u.FirstName == "" {
			return errors.New("required firstname")
		}
		if u.LastName == "" {
			return errors.New("required lastname")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	}
}

func (u *User) SaveUser(db gorm.DB) (*User, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserById(db *gorm.DB, uid uint32) (*User, error) {
	err := db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &User{}, errors.New("user not found")
	}
	return u, err
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	err := u.BeforeSave(db)
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"email":      u.Email,
			"password":   u.Password,
			"first_name": u.FirstName,
			"last_name":  u.LastName,
			"update_at":  time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
