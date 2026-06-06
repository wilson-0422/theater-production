package services

import (
	"theater-production/src/config"
	"theater-production/src/models"
)

func GetAllTours() ([]models.Tour, error) {
	var tours []models.Tour
	err := config.GetDB().Preload("Script").Order("created_at DESC").Find(&tours).Error
	return tours, err
}

func GetTourByID(id uint) (*models.Tour, error) {
	var tour models.Tour
	err := config.GetDB().Preload("Script").First(&tour, id).Error
	return &tour, err
}

func CreateTour(tour *models.Tour) error {
	return config.GetDB().Create(tour).Error
}

func UpdateTour(tour *models.Tour) error {
	return config.GetDB().Save(tour).Error
}

func DeleteTour(id uint) error {
	return config.GetDB().Delete(&models.Tour{}, id).Error
}

func CountTours() (int64, error) {
	var count int64
	err := config.GetDB().Model(&models.Tour{}).Count(&count).Error
	return count, err
}

func GetToursByStatus(status string) ([]models.Tour, error) {
	var tours []models.Tour
	err := config.GetDB().Where("status = ?", status).Preload("Script").Order("start_date ASC").Find(&tours).Error
	return tours, err
}
