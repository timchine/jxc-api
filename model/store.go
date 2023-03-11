package model

type CargoStore struct {
	Model
	CsID          int     `json:"cs_id" gorm:"primaryKey;autoIncrement"`
	CargoID       int     `json:"cargo_id"`                        //关联货品
	Surplus       float64 `json:"surplus" gorm:"default:0"`        //余量
	PreOut        float64 `json:"pre_out" gorm:"default:0"`        //准备出库
	Usable        float64 `json:"usable" gorm:"default:0"`         //可用余量
	PrePut        float64 `json:"pre_put" gorm:"default:0"`        //准备入库
	UpWarning     float64 `json:"up_warning" gorm:"default:0"`     //可用余量告警
	LowWarning    float64 `json:"low_warning" gorm:"default:0"`    //可用余量告警
	WarningStatus int     `json:"warning_status" gorm:"default:0"` // 1 开启上限 2 开启下限
}
