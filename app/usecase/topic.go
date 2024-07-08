package usecase

import (
	"errors"

	"github.com/zakirkun/zot-skill-test/app/domain/models"
	"github.com/zakirkun/zot-skill-test/app/domain/types"
	"github.com/zakirkun/zot-skill-test/app/repository"
	"github.com/zakirkun/zot-skill-test/utils"
)

func GetDetailTopic(id int) (*models.Topic, error) {

	topic, err := repository.GetDetailTopic(id)
	if err != nil {
		return nil, err
	}

	if topic.ID == 0 {
		return nil, errors.New("topic not found")
	}

	return topic, nil
}

func CreateTopic(request types.CreateTopic) (bool, error) {
	err := repository.CreateTopic(models.Topic{
		Name: request.Name,
		Slug: utils.Slugify(request.Name),
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func ListTopic() (*[]models.Topic, error) {
	topics, err := repository.GetAllTopic()
	if err != nil {
		return nil, err
	}

	if len(*topics) == 0 {
		return nil, errors.New("topic empty")
	}

	return topics, nil
}

func SearchTopic(keyword string) (*[]models.Topic, error) {

	cond := map[string]interface{}{
		"name LIKE ?": keyword,
	}

	topics, err := repository.GetAllTopicWithCond(cond)
	if err != nil {
		return nil, err
	}

	if len(*topics) == 0 {
		return nil, errors.New("topic empty")
	}

	return topics, nil
}

func UpdateTopic(id int, request types.UpdateeTopic) (bool, error) {

	topic, err := repository.GetDetailTopic(id)
	if err != nil {
		return false, err
	}

	if topic.ID == 0 {
		return false, errors.New("topic not found")
	}

	err = repository.UpdateTopic(id, models.Topic{
		Name: request.Name,
		Slug: utils.Slugify(request.Name),
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteTopic(id int) (bool, error) {
	err := repository.DeleteTopic(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
