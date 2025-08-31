package models

import "time"

type Visit struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	URLID     uint      `gorm:"not null;index" json:"url_id"`
	IPAddress string    `gorm:"not null" json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	Referrer  string    `json:"referrer"`
	VisitTime time.Time `json:"visit_time"`
	CreatedAt time.Time `json:"created_at"`
}