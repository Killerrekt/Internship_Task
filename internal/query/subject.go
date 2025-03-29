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

func CheckSubjects(subs []string) bool {
	var count int64
	err := database.DB.Model(&model.Subject{}).Where("name in ?", subs).Count(&count).Error
	if err != nil {
		return false
	} else {
		return count == int64(len(subs))
	}
}
