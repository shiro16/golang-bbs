package models

import (
	"time"
)

type Comment struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	Nickname  string     `sql:"size:64" json:"nickname" validate:"max=64"`
	Body      string     `sql:"size:255" json:"body" validate:"min=1,max=255"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
