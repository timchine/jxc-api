package router

import (
	"github.com/gin-gonic/gin"
	"github.com/timchine/jxc/api"
	"gorm.io/gorm"
)

func cargoRouter(r *gin.RouterGroup, db *gorm.DB) {
	a := api.NewCargoApi(db)
	r.POST("cargo_kind", a.AddCargoKind())
}
