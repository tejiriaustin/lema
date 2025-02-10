package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	Models interface {
		GetID() uuid.UUID
		GetVersion() uint
	}

	PreValidator interface {
		PreValidate()
	}

	AccountInfo struct {
		Id       string `json:"id"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
	}
)

type Shared struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"deleted_at"`
	Version   uint       `json:"version" gorm:"_version"`
}

func (m Shared) PreValidate() {}

var _ Models = &Shared{}

func (m Shared) GetID() uuid.UUID {
	return m.ID
}

func (m Shared) GetVersion() uint {
	return m.Version
}

func (m Shared) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}
