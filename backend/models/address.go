package models

import (
	"fmt"
	"github.com/google/uuid"
)

type Address struct {
	Shared  `gorm:"embedded"`
	UserID  string `json:"user_id" gorm:"type:varchar(32);not null"`
	Street  string `json:"street" gorm:"type:varchar(100);not null"`
	City    string `json:"city" gorm:"type:varchar(100);not null"`
	State   string `json:"state" gorm:"type:varchar(100);not null"`
	Zipcode string `json:"zipcode" gorm:"type:varchar(20);not null"`
}

func (a *Address) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.Zipcode)
}

func (a *Address) PreValidate() {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
}
