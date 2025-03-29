package query

import (
	"github.com/killerrekt/quantstrategix/internal/database"
	"github.com/killerrekt/quantstrategix/internal/model"
)

func AddSubject(incoming model.AddSubjectReq) error {
	return database.DB.Model(&model.Subject{}).Create(&model.Subject{Name: incoming.Name}).Error
}

func GetAllSubject() ([]model.Subject, error) {
	var data []model.Subject
	err := database.DB.Model(&model.Subject{}).Find(&data).Error
	return data, err
}
