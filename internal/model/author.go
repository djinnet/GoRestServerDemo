package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primary_key;type:uuid"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
}
