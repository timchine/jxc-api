package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Model struct {
	CreatedAt LocalTime `json:"created_at" gorm:"column:created_at;autoCreateTime;type:timestamp"`
	UpdatedAt LocalTime `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;type:timestamp"`
}

type LocalTime int64

func (t LocalTime) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", time.Unix(int64(t), 0).Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	tlt := time.Unix(int64(t), 0)
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value.Unix())
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
