package model

type Channel struct {
	Model
	ChannelID   int    `json:"channel_id" gorm:"primaryKey;autoIncrement;column:channel_id;type:int;comment:主键"`
	ChannelName string `json:"channel_name" gorm:"column:channel_name;type:varchar(100);comment:供应商名称"`
	Intro       string `json:"intro"`
	Address     string `json:"address"`
}
