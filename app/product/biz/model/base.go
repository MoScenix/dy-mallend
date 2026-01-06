package model

import "time"

type Base struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
