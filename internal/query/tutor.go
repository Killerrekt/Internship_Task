package query

import (
	"github.com/killerrekt/quantstrategix/internal/database"
	"github.com/killerrekt/quantstrategix/internal/model"
)

func GetTutor(email string) (model.Tutor, error) {
	var data model.Tutor
	subquery := database.DB.Model(&model.User{}).Where("email = ?", email).Select("id")

	err := database.DB.Model(&model.Tutor{}).
		Where("user_id IN (?)", subquery).
		First(&data).Error

	return data, err
}

func AddDetail(incoming model.Tutor) error {
	return database.DB.Model(&model.Tutor{}).Where("user_id = ?", incoming.UserID).Save(&incoming).Error
}

func CreateDetail(incoming model.Tutor) error {
	return database.DB.Model(&model.Tutor{}).Create(&incoming).Error
}

func GetTutorBySubject(subject string) ([]model.Tutor, error) {
	var data []model.Tutor
	err := database.DB.Model(&model.Tutor{}).Where("? = ANY(tutors.subjects_name)", subject).Find(&data).Error
	return data, err
}
