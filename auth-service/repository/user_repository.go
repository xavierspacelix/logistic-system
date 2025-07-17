package repository

import (
	"github.com/xavierspacelix/logistic-system/auth-service/config"
	"github.com/xavierspacelix/logistic-system/auth-service/model"
)

func FindUserByMSISDN(msisdn string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("msisdn = ?", msisdn).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByUsernameOrMSISDN(username, msisdn string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("username = ? OR msisdn = ?", username, msisdn).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByID(id string) (*model.User, error) {
	var user model.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *model.User) error {
	return config.DB.Create(user).Error
}
