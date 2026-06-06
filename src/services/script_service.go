package services

import (
	"theater-production/src/config"
	"theater-production/src/models"
)

func GetAllScripts() ([]models.Script, error) {
	var scripts []models.Script
	err := config.GetDB().Order("created_at DESC").Find(&scripts).Error
	return scripts, err
}

func GetScriptByID(id uint) (*models.Script, error) {
	var script models.Script
	err := config.GetDB().First(&script, id).Error
	return &script, err
}

func CreateScript(script *models.Script) error {
	return config.GetDB().Create(script).Error
}

func UpdateScript(script *models.Script) error {
	return config.GetDB().Save(script).Error
}

func DeleteScript(id uint) error {
	return config.GetDB().Delete(&models.Script{}, id).Error
}

func GetScriptsByStatus(status string) ([]models.Script, error) {
	var scripts []models.Script
	err := config.GetDB().Where("status = ?", status).Order("created_at DESC").Find(&scripts).Error
	return scripts, err
}

func CountScripts() (int64, error) {
	var count int64
	err := config.GetDB().Model(&models.Script{}).Count(&count).Error
	return count, err
}
