package router

import (
	"github.com/gin-gonic/gin"
	"github.com/timchine/jxc/api"
	"gorm.io/gorm"
)

func cargoRouter(r *gin.RouterGroup, db *gorm.DB) {
	a := api.NewCargoApi(db)
	//货品种类
	r.POST("cargo_kind", a.AddCargoKind())
	r.GET("cargo_kind/:ck_id", a.GetCargoKind())
	r.PUT("cargo_kind", a.UpdateCargoKind())
	r.DELETE("cargo_kind/:ck_id", a.DeleteCargoKind())
	r.GET("cargo_kind/_search", a.SearchCargoKind())

	r.POST("cargo", a.AddCargo())
	r.DELETE("cargo/:cargo_id", a.DeleteCargo())
	r.PUT("cargo", a.UpdateCargo())
	r.GET("cargo/:cargo_id", a.GetCargo())
	r.GET("cargo/_search", a.SearchCargo())

	//上传图片 会生成一张缩略图一张原图
	r.POST("/image", a.UploadImage())
}
