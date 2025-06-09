package models

import (
	"time"
)

type Tasks struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null" json:"created_at"`
}
