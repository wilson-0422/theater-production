package services

import (
	"theater-production/src/config"
	"theater-production/src/models"
)

func GetAllProps() ([]models.Prop, error) {
	var props []models.Prop
	err := config.GetDB().Order("created_at DESC").Find(&props).Error
	return props, err
}

func GetPropByID(id uint) (*models.Prop, error) {
	var prop models.Prop
	err := config.GetDB().First(&prop, id).Error
	return &prop, err
}

func CreateProp(prop *models.Prop) error {
	return config.GetDB().Create(prop).Error
}

func UpdateProp(prop *models.Prop) error {
	return config.GetDB().Save(prop).Error
}

func DeleteProp(id uint) error {
	return config.GetDB().Delete(&models.Prop{}, id).Error
}

func GetAllPropRequisitions() ([]models.PropRequisition, error) {
	var requisitions []models.PropRequisition
	err := config.GetDB().Preload("Prop").Preload("Actor").Order("created_at DESC").Find(&requisitions).Error
	return requisitions, err
}

func GetPropRequisitionsByPropID(propID uint) ([]models.PropRequisition, error) {
	var requisitions []models.PropRequisition
	err := config.GetDB().Where("prop_id = ?", propID).Preload("Actor").Order("created_at DESC").Find(&requisitions).Error
	return requisitions, err
}

func CreatePropRequisition(req *models.PropRequisition) error {
	return config.GetDB().Create(req).Error
}

func UpdatePropRequisition(req *models.PropRequisition) error {
	return config.GetDB().Save(req).Error
}

func CountProps() (int64, error) {
	var count int64
	err := config.GetDB().Model(&models.Prop{}).Count(&count).Error
	return count, err
}
