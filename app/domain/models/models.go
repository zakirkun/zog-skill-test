package models

import "time"

type StatusNews string

const (
	Draft     StatusNews = "draft"
	Deleted   StatusNews = "deleted"
	Published StatusNews = "published"
)

type Topic struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	Slug      string     `gorm:"type:varchar(255);not null" json:"slug"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	News      []News     `gorm:"many2many:news_topics;" json:"news"`
}

type News struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"type:varchar(255);not null" json:"title"`
	Slug      string     `gorm:"type:varchar(255);not null" json:"slug"`
	Thumbnail string     `gorm:"type:text;not null" json:"thumbnail"`
	Status    StatusNews `gorm:"type:status_news;default:'draft'" json:"status"`
	Content   string     `gorm:"type:text" json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Topics    []Topic    `gorm:"many2many:news_topics;" json:"topics"`
}

type NewsTopic struct {
	ID      uint `gorm:"primaryKey"`
	NewsID  uint `gorm:"index"`
	TopicID uint `gorm:"index"`
}
