package model

import "github.com/google/uuid"

type Subject struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string    `gorm:"type:text;not null"`
}
