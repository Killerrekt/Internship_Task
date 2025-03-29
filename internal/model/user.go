package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string    `gorm:"type:text;not null"`
	Email    string    `gorm:"type:text;unique;not null"`
	Password string    `gorm:"type:text;not null"`
	Role     string    `gorm:"type:text;not null"`
}
