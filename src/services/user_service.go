package services

import (
	"theater-production/src/config"
	"theater-production/src/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password, role string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}
	return config.GetDB().Create(&user).Error
}

func AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	if err := config.GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := config.GetDB().First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.GetDB().Order("created_at DESC").Find(&users).Error
	return users, err
}
