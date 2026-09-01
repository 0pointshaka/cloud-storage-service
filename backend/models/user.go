package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Files    []File `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = 0
	u.UpdatedAt = 0
	return nil
}
