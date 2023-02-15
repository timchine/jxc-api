package model

type Provider struct {
	Model
	ProviderID        int       `json:"provider_id" gorm:"primaryKey;autoIncrement;column:provider_id;type:int;comment:主键"`
	ProviderShortName string    `json:"provider_short_name" gorm:"column:provider_short_name;type:varchar(50);comment:供应商简称"`
	Intro             string    `json:"intro" gorm:"column:intro;type:varchar(500);comment:简介"`
	Contact           string    `json:"contact" gorm:"column:contact;type:varchar(50);comment:联系人"`
	ContactPhone      string    `json:"contact_phone" gorm:"column:contact_phone;type:varchar(20);comment:联系人电话"`
	ProviderName      string    `json:"provider_name" gorm:"column:provider_name;type:varchar(50);comment:供应商全称"`
	TaxNumber         string    `json:"tax_number" gorm:"column:tax_number;type:varchar(50);comment:税号"`
	Address           string    `json:"address" gorm:"column:address;type:varchar(300);comment:地址"` //单位地址
	ProviderTel       string    `json:"provider_tel" gorm:"column:provider_tel;type:varchar(20);comment:单位电话"`
	CardNo            string    `json:"card_no" gorm:"column:card_no;type:varchar(50);comment:银行卡号"`
	OpeningBank       string    `json:"opening_bank" gorm:"column:opening_bank;type:varchar(30);comment:开户行"`
	PayMethod         PayMethod `json:"pay_method" gorm:"column:pay_method;type:int;comment:付款方式"`
	Status            int       `json:"status" gorm:"column:status;type:int;comment:状态 1 正常 8 删除；default:1"`
}

// todo 供应商对应多个原材料
type ProviderCargo struct {
	Model
	PcID       int     `json:"pc_id"`
	ProviderID int     `json:"provider_id"`
	CargoID    int     `json:"cargo_id"`
	Price      float64 `json:"price"` //金额浮点数问题
	MeasureID  int     `json:"measure_id"`
}
