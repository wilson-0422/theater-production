package services

import (
	"theater-production/src/config"
	"theater-production/src/models"
)

func GetAllActors() ([]models.Actor, error) {
	var actors []models.Actor
	err := config.GetDB().Order("created_at DESC").Find(&actors).Error
	return actors, err
}

func GetActorByID(id uint) (*models.Actor, error) {
	var actor models.Actor
	err := config.GetDB().First(&actor, id).Error
	return &actor, err
}

func CreateActor(actor *models.Actor) error {
	return config.GetDB().Create(actor).Error
}

func UpdateActor(actor *models.Actor) error {
	return config.GetDB().Save(actor).Error
}

func DeleteActor(id uint) error {
	return config.GetDB().Delete(&models.Actor{}, id).Error
}

func GetActorSchedules(actorID uint) ([]models.ActorSchedule, error) {
	var schedules []models.ActorSchedule
	err := config.GetDB().Where("actor_id = ?", actorID).Preload("Script").Order("schedule_date DESC").Find(&schedules).Error
	return schedules, err
}

func GetAllActorSchedules() ([]models.ActorSchedule, error) {
	var schedules []models.ActorSchedule
	err := config.GetDB().Preload("Actor").Preload("Script").Order("schedule_date DESC").Find(&schedules).Error
	return schedules, err
}

func CreateActorSchedule(schedule *models.ActorSchedule) error {
	return config.GetDB().Create(schedule).Error
}

func CountActors() (int64, error) {
	var count int64
	err := config.GetDB().Model(&models.Actor{}).Count(&count).Error
	return count, err
}
