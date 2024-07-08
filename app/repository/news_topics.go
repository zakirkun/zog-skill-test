package repository

import (
	"github.com/zakirkun/zot-skill-test/app/domain/models"
	"github.com/zakirkun/zot-skill-test/pkg/database"
)

func CreateNewsTopic(data models.NewsTopic) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func DeleteNewsTopic(topic int, news int) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	var data models.NewsTopic
	if err := db.Model(&models.NewsTopic{}).Where("news_id = ? and topic_id = ?", news, topic).Find(&data).Error; err != nil {
		return err
	}

	if err := db.Delete(&data).Error; err != nil {
		return err
	}

	return nil
}
