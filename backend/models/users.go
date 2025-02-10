package models

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Shared   `gorm:"embedded"`
	FullName string   `json:"full_name" gorm:"type:varchar(100);not null"`
	Email    string   `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Address  *Address `json:"address" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Posts    []Post   `json:"posts,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (u *User) PreValidate() {
	fmt.Println("called prevalidate")
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
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
		if u.Address.ID == uuid.Nil {
			u.Address.ID = uuid.New()
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
