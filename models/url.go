package models

import (
	"time"
)

type URL struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	OriginalURL  string    `gorm:"not null" json:"original_url"`
	ShortToken   string    `gorm:"uniqueIndex;not null" json:"short_token"`
	ClickCount   int       `gorm:"default:0" json:"click_count"`
	UniqueVisits int       `gorm:"default:0" json:"unique_visits"`
	Visits       []Visit   `json:"-" gorm:"foreignKey:URLID"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}