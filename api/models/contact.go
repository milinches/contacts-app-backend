package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type (
	Contact struct {
		ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
		Email       string         `gorm:"type:string" json:"email"`
		Name        string         `gorm:"type:string;not null" json:"name"`
		Address     string         `gorm:"type:text;" json:"address"`
		PhoneNumber pq.StringArray `gorm:"type:text[];not null" json:"phone_number"`
		Type        pq.StringArray `gorm:"type:text[];not null" json:"type"`
		User        User           `json:"user"`
		UserID      uint           `json:"user_id"`
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
		DeletedAt   time.Time      `gorm:"index" json:"deleted_at"`
	}
)

// Added Hooks — Gorm hooks are function that are called before or after creation

func (u *Contact) TableName() string {
	return "contact"
}

func (c *Contact) Prepare() {
	c.ID = 0
	c.Email = html.EscapeString(strings.TrimSpace(c.Email))
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.Address = html.EscapeString(strings.TrimSpace(c.Address))
	c.PhoneNumber = []string{}
	c.Type = []string{}
	c.User = User{}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Contact) Validate() error {
	if c.Name == "" {
		return errors.New("required field")
	}
	if len(c.PhoneNumber) == 0 {
		return errors.New("required field")
	}
	if len(c.Type) == 0 {
		return errors.New("required field")
	}
	if c.UserID < 1 {
		return errors.New("required user")
	}
	return nil
}

func (c *Contact) SaveContact(db *gorm.DB) (*Contact, error) {
	var err error
	if err = db.Debug().Model(&Contact{}).Create(&c).Error; err != nil {
		return &Contact{}, err
	}
	if c.ID != 0 {
		if err = db.Debug().Model(&Contact{}).Where("id = ?", c.UserID).Take(&c.User).Error; err != nil {
			return &Contact{}, err
		}
	}
	return c, nil
}

func (c *Contact) FindAllContacts(db *gorm.DB) (*[]Contact, error) {
	var err error
	contacts := []Contact{}
	if err = db.Debug().Model(&Contact{}).Limit(100).Find(&contacts).Error; err != nil {
		return &[]Contact{}, err
	}
	if len(contacts) > 0 {
		for i := range contacts {
			if err = db.Debug().Model(&Contact{}).Where("id = ?", contacts[i].UserID).Take(&contacts[i].User).Error; err != nil {
				return &[]Contact{}, err
			}
		}
	}
	return &contacts, nil
}

func (c *Contact) FindContactByID(db *gorm.DB, id uint64) (*Contact, error) {
	var err error
	if err = db.Debug().Model(&Contact{}).Where("id = ?", id).Take(&c).Error; err != nil {
		return &Contact{}, err
	}
	if c.ID != 0 {
		if err = db.Debug().Model(&User{}).Where("id = ?", c.UserID).Take(&c.User).Error; err != nil {
			return &Contact{}, nil
		}
	}
	return c, nil
}

func (c *Contact) DeleteContact(db *gorm.DB, cid uint64, uid uint32) (int64, error) {
	db = db.Debug().Model(&Contact{}).Where("id = ? and user_id = ?", cid, uid).Take(&Contact{}).Delete(&Contact{})
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("contact not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
