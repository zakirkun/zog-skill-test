package usecase

import (
	"errors"

	"github.com/zakirkun/zot-skill-test/app/domain/models"
	"github.com/zakirkun/zot-skill-test/app/domain/types"
	"github.com/zakirkun/zot-skill-test/app/repository"
	"github.com/zakirkun/zot-skill-test/utils"
)

func UpdateNews(id int, request types.UpdateNews) (bool, error) {

	news, err := repository.GetDetailNews(id)
	if err != nil {
		return false, err
	}

	if news.ID == 0 {
		return false, errors.New("news not found")
	}

	insert := models.News{
		Title:   request.Title,
		Slug:    utils.Slugify(request.Title),
		Status:  models.StatusNews(request.Status),
		Content: request.Content,
	}

	if request.URLThumbnail != "" {
		insert.Thumbnail = request.URLThumbnail
	}

	if err := repository.UpdateNews(id, insert); err != nil {
		return false, err
	}

	for _, topic := range request.Topics {
		repository.DeleteNewsTopic(topic, id)
	}

	for _, topic := range request.Topics {
		_ = repository.CreateNewsTopic(models.NewsTopic{TopicID: uint(topic), NewsID: insert.ID})
	}

	return true, nil
}

func CreateNews(request types.CreateNews) (bool, error) {

	insert := models.News{
		Title:     request.Title,
		Slug:      utils.Slugify(request.Title),
		Status:    models.StatusNews(request.Status),
		Content:   request.Content,
		Thumbnail: request.URLThumbnail,
	}

	if err := repository.CreateNews(insert); err != nil {
		return false, err
	}

	for _, topic := range request.Topics {
		_ = repository.CreateNewsTopic(models.NewsTopic{TopicID: uint(topic), NewsID: insert.ID})
	}

	return true, nil
}

func GetDetailNews(id int) (*models.News, error) {

	news, err := repository.GetDetailNews(id)
	if err != nil {
		return nil, err
	}

	if news.ID == 0 {
		return nil, errors.New("news not found")
	}
	return news, nil
}

func ListNews() (*[]models.News, error) {
	news, err := repository.GetAllNews()
	if err != nil {
		return nil, err
	}

	if len(*news) == 0 {
		return nil, errors.New("news empty")
	}

	return news, nil
}

func DeleteNews(id int) (bool, error) {
	err := repository.DeleteNews(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func FilterNews(status string) (*[]models.News, error) {
	news, err := repository.GetAllNewsWithCond(map[string]interface{}{
		"status": models.StatusNews(status),
	})
	if err != nil {
		return nil, err
	}

	if len(*news) == 0 {
		return nil, errors.New("news empty")
	}

	return news, nil
}
