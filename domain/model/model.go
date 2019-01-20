package model

import (
	"time"

	"github.com/rs/xid"
)

// Model is common table model.
type Model struct {
	ID        string     `json:"id" gorm:"primary_key;size:20"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// BeforeCreate set id.
func (m *Model) BeforeCreate() (err error) {
	if len(m.ID) == 0 {
		m.ID = xid.New().String()
	}
	return nil
}
