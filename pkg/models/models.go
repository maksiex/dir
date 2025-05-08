package models

import (
	"gorm.io/gorm"
	"time"
)

type Flight struct {
	gorm.Model
	FlightNumber string
	Departure    string
	Arrival      string
	Status       string
}

type NewsFromMediastack struct {
	ID          uint      `gorm:"primaryKey" json:"-"`
	Author      *string   `json:"author"` // or null
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Source      string    `json:"source"`
	Image       *string   `json:"image"` // or null
	Category    string    `json:"category"`
	Language    string    `json:"language"`
	Country     string    `json:"country"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time
}

type MediastackResponse struct {
	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
		Total  int `json:"total"`
	} `json:"pagination"`
	Data []NewsFromMediastack `json:"data"`
}

type NewsFromCurrents struct {
	ID          string    `gorm:"primaryKey" json:"id"` // UUID
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Author      string    `json:"author"`
	Image       string    `json:"image"`
	Language    string    `json:"language"`
	Category    string    `json:"-"`
	Published   time.Time `json:"published"`
	CreatedAt   time.Time
}

type CurrentsAPIResponse struct {
	Status string               `json:"status"`
	News   []RawCurrentsArticle `json:"news"`
}

type RawCurrentsArticle struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Author      string   `json:"author"`
	Image       string   `json:"image"`
	Language    string   `json:"language"`
	Category    []string `json:"category"`
	Published   string   `json:"published"`
}

type NewsFromGnews struct {
	ID          uint      `gorm:"primaryKey" json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	URL         string    `json:"url"`
	Image       string    `json:"image"`
	PublishedAt time.Time `json:"publishedAt"`
	SourceName  string    `json:"-"`
	SourceURL   string    `json:"-"`
	CreatedAt   time.Time
}

type GNewsAPIResponse struct {
	TotalArticles int               `json:"totalArticles"`
	Articles      []RawGNewsArticle `json:"articles"`
}

type RawGNewsArticle struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	URL         string    `json:"url"`
	Image       string    `json:"image"`
	PublishedAt time.Time `json:"publishedAt"`
	Source      struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"source"`
}

type NewsFromNewsapi struct {
	ID          uint      `gorm:"primaryKey" json:"-"`
	SourceID    *string   `json:"-"`
	SourceName  string    `json:"-"`
	Author      *string   `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
	CreatedAt   time.Time
}

type NewsAPIResponse struct {
	Status       string              `json:"status"`
	TotalResults int                 `json:"totalResults"`
	Articles     []RawNewsapiArticle `json:"articles"`
}

type RawNewsapiArticle struct {
	Source struct {
		ID   *string `json:"id"`
		Name string  `json:"name"`
	} `json:"source"`
	Author      *string   `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}
