package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Shared   `gorm:"embedded"`
	Name     string   `json:"name" gorm:"type:varchar(200);not null"`
	Username string   `json:"username" gorm:"type:varchar(100);not null"`
	Email    string   `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Address  *Address `json:"address" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Posts    []Post   `json:"posts,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (u *User) PreValidate() {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	if u.CreatedAt == nil {
		now := time.Now().UTC()
		u.CreatedAt = &now
	}

	if u.Version > 0 {
		u.Version++
	} else {
		u.Version = 1
	}

	if u.Address != nil {
		if u.Address.ID == "" {
			u.Address.ID = uuid.New().String()
		}

		if u.Address.CreatedAt == nil {
			now := time.Now().UTC()
			u.Address.CreatedAt = &now
		}

		if u.Address.Version > 0 {
			u.Address.Version++
		} else {
			u.Address.Version = 1
		}
	}
}
