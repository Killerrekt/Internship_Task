package query

import (
	"github.com/killerrekt/quantstrategix/internal/database"
	"github.com/killerrekt/quantstrategix/internal/model"
	"gorm.io/gorm/clause"
)

func CreateBooking(data model.Booking) error {
	return database.DB.Model(&model.Booking{}).Create(&data).Error
}

func GetBooking(email string) ([]model.Booking, error) {
	var data []model.Booking
	subquery := database.DB.Model(&model.User{}).Where("email = ?", email).Select("id")
	err := database.DB.Model(&model.Booking{}).Preload(clause.Associations).Where("student_id = (?)", subquery).Find(&data).Error
	return data, err
}

func UpdateBooking(req model.UpdateBookingReq) error {
	return database.DB.Model(&model.Booking{}).Where("id = ?", req.Id).Update("status", req.Status).Error
}
