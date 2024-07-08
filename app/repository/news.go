package repository

import (
	"github.com/zakirkun/zot-skill-test/app/domain/models"
	"github.com/zakirkun/zot-skill-test/pkg/database"
)

func CreateNews(data models.News) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func UpdateNews(id int, data models.News) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Model(&models.News{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func GetAllNews() (*[]models.News, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []models.News
	if err := db.Debug().Model(&models.News{}).Preload("Topics").Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func DeleteNews(id int) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Delete(models.News{}, id).Error; err != nil {
		return err
	}

	return nil
}

func GetDetailNews(id int) (*models.News, error) {

	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.News
	if err := db.Model(&models.News{}).Preload("Topics").Where("id = ?", id).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func GetAllNewsWithCond(cond map[string]interface{}) (*[]models.News, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []models.News
	if err := db.Debug().Model(&models.News{}).Preload("Topics").Where(cond).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
