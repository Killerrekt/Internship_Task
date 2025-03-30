package model

import "github.com/google/uuid"

type RegisterReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type LoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AddSubjectReq struct {
	Name string `json:"name" binding:"required"`
}

type AddTutorDetail struct {
	Bio         string   `json:"bio"`
	SubjectName []string `json:"subject_names"`
}

type Bookingreq struct {
	TutorId     uuid.UUID `json:"tutor_id" binding:"required"`
	Subject     uuid.UUID `json:"subject_id" binding:"required"`
	BookingTime string    `json:"booking_time" binding:"required"`
}

type UpdateBookingReq struct {
	Id     string `uri:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}
