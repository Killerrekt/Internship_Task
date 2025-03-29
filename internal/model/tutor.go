package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Tutor struct {
	UserID       uuid.UUID      `gorm:"type:uuid;not null;primary_key"`
	Bio          string         `gorm:"type:text"`
	SubjectsName pq.StringArray `gorm:"type:text[]"`
}
