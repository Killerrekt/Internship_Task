package model

import (
	"github.com/google/uuid"
)

type Booking struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	StudentID   uuid.UUID `gorm:"type:uuid"`
	Student     User      `gorm:"foreignKey:StudentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TutorID     uuid.UUID `gorm:"type:uuid"`
	Tutor       User      `gorm:"foreignKey:TutorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SubjectID   uuid.UUID `gorm:"type:uuid"`
	Subject     Subject   `gorm:"foreignKey:SubjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BookingTime string
	Status      string `gorm:"text"`
}
