package model

import (
	"time"
)

type Model struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime;type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:timestamp"`
}
