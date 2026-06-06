package services

import (
	"theater-production/src/config"
	"theater-production/src/models"
)

func GetAllTheaters() ([]models.Theater, error) {
	var theaters []models.Theater
	err := config.GetDB().Order("created_at DESC").Find(&theaters).Error
	return theaters, err
}

func GetTheaterByID(id uint) (*models.Theater, error) {
	var theater models.Theater
	err := config.GetDB().First(&theater, id).Error
	return &theater, err
}

func CreateTheater(theater *models.Theater) error {
	return config.GetDB().Create(theater).Error
}

func UpdateTheater(theater *models.Theater) error {
	return config.GetDB().Save(theater).Error
}

func DeleteTheater(id uint) error {
	return config.GetDB().Delete(&models.Theater{}, id).Error
}

func GetAllTheaterSchedules() ([]models.TheaterSchedule, error) {
	var schedules []models.TheaterSchedule
	err := config.GetDB().Preload("Theater").Preload("Script").Order("start_time DESC").Find(&schedules).Error
	return schedules, err
}

func GetTheaterSchedulesByTheaterID(theaterID uint) ([]models.TheaterSchedule, error) {
	var schedules []models.TheaterSchedule
	err := config.GetDB().Where("theater_id = ?", theaterID).Preload("Script").Order("start_time DESC").Find(&schedules).Error
	return schedules, err
}

func CreateTheaterSchedule(schedule *models.TheaterSchedule) error {
	return config.GetDB().Create(schedule).Error
}

func CountTheaters() (int64, error) {
	var count int64
	err := config.GetDB().Model(&models.Theater{}).Count(&count).Error
	return count, err
}
