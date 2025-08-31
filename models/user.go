package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name  string    	`gorm:"uniqueIndex;not null" json:"name"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}