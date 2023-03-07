package model

type CargoStore struct {
	Model
	CkID    int     `json:"ck_id" gorm:"primaryKey;autoIncrement"`
	CargoID int     `json:"cargo_id"`
	Surplus float64 `json:"surplus"` //余量
}
