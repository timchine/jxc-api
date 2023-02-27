package model

type Image struct {
	Model
	ImageID       int    `json:"image_id" gorm:"primaryKey;autoIncrement"`
	ThumbnailName string `json:"thumbnail_name"` //缩略图
	ImageName     string `json:"image_name"`
	ImageHash     string `json:"image_hash"`
	Status        int    `json:"status" gorm:"default:1"` //1 未使用 2 被使用 8 删除
}
