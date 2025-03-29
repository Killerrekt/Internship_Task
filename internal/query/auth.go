package query

import (
	"github.com/killerrekt/quantstrategix/internal/database"
	"github.com/killerrekt/quantstrategix/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func SaveUser(incoming model.RegisterReq) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(incoming.Password), 10)
	incoming.Password = string(hashedPassword)

	err := database.DB.Model(&model.User{}).Create(&model.User{
		Name:     incoming.Name,
		Email:    incoming.Email,
		Password: incoming.Password,
		Role:     incoming.Role,
	}).Error

	return err
}

func FindUser(email string) (model.User, error) {
	var user model.User
	err := database.DB.Model(&model.User{}).Where("email = ?", email).First(&user).Error
	return user, err
}
