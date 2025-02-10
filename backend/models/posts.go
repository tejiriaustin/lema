package models

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	Shared `gorm:"embedded"`
	UserID string `json:"user_id" gorm:"not null"`
	Title  string `json:"title" gorm:"type:varchar(200);not null"`
	Body   string `json:"body" gorm:"type:text;not null"`
}

func (p *Post) PreValidate() {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}

	if p.CreatedAt == nil {
		now := time.Now().UTC()
		p.CreatedAt = &now
	}

	if p.Version > 0 {
		p.Version++
	} else {
		p.Version = 1
	}

	p.Version = 1
}
