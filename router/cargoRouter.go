package router

import (
	"github.com/gin-gonic/gin"
	"github.com/timchine/jxc/api"
	"gorm.io/gorm"
)

func cargoRouter(r *gin.RouterGroup, db *gorm.DB) {
	a := api.NewCargoApi(db)
	r.POST("cargo_kind", a.AddCargoKind())
	r.GET("cargo_kind/:ck_id", a.GetCargoKind())
	r.PUT("cargo_kind", a.UpdateCargoKind())
	r.DELETE("cargo_kind/:ck_id", a.DeleteCargoKind())
	r.GET("cargo_kind/_search", a.SearchCargoKind())
}
