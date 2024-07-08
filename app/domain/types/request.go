package types

import (
	"mime/multipart"

	"github.com/zakirkun/zot-skill-test/app/domain/models"
)

type CreateTopic struct {
	Name string `json:"name"`
}

type UpdateeTopic struct {
	Name string `json:"name"`
}

type CreateNews struct {
	Title        string                `form:"title"`
	Status       models.StatusNews     `form:"status"`
	Content      string                `form:"content"`
	Topics       []int                 `form:"topics"`
	Thumbnail    *multipart.FileHeader `form:"thumbnail"`
	URLThumbnail string
}

type UpdateNews struct {
	Title        string                `form:"title"`
	Slug         string                `form:"slug"`
	Status       models.StatusNews     `form:"status"`
	Content      string                `form:"content"`
	Topics       []int                 `form:"topics"`
	Thumbnail    *multipart.FileHeader `form:"thumbnail"`
	URLThumbnail string
}
