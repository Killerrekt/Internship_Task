package model

import "github.com/google/uuid"

type RegisterReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddSubjectReq struct {
	Name string `json:"name"`
}

type AddTutorDetail struct {
	Bio         string   `json:"bio"`
	SubjectName []string `json:"subject_names"`
}

type Bookingreq struct {
	TutorId     uuid.UUID `json:"tutor_id"`
	Subject     uuid.UUID `json:"subject_id"`
	BookingTime string    `json:"booking_time"`
}

type UpdateBookingReq struct {
	Id     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}
