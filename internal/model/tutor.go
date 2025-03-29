package model

import "github.com/google/uuid"

type Tutor struct {
	UserID     uuid.UUID   `gorm:"type:uuid;primaryKey"`
	Name       string      `gorm:"type:text;not null"`
	Subjects   []Subject   `gorm:"many2many:tutor_subjects;"`
	SubjectIDs []uuid.UUID `gorm:"-"`
}
