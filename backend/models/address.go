package models

import (
	"fmt"
	"github.com/google/uuid"
)

type Address struct {
	Shared  `gorm:"embedded"`
	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Street  string    `json:"street" gorm:"type:varchar(200);not null"`
	City    string    `json:"city" gorm:"type:varchar(100);not null"`
	State   string    `json:"state" gorm:"type:varchar(100);not null"`
	ZipCode string    `json:"zipcode" gorm:"type:varchar(20);not null"`
}

func (a *Address) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.ZipCode)
}

func (a *Address) PreValidate() {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
}
