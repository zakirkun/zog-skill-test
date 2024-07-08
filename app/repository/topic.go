package repository

import (
	"github.com/zakirkun/zot-skill-test/app/domain/models"
	"github.com/zakirkun/zot-skill-test/pkg/database"
)

func GetDetailTopic(id int) (*models.Topic, error) {

	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.Topic
	if err := db.Preload("News").Model(&models.Topic{}).Where("id = ?", id).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func GetAllTopicWithCond(cond map[string]interface{}) (*[]models.Topic, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []models.Topic
	if err := db.Debug().Model(&models.Topic{}).Where(cond).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func GetAllTopic() (*[]models.Topic, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []models.Topic
	if err := db.Model(&models.Topic{}).Preload("News").Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func CreateTopic(data models.Topic) error {

	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTopic(id int, data models.Topic) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Model(&models.Topic{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTopic(id int) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Delete(models.Topic{}, id).Error; err != nil {
		return err
	}

	return nil
}
