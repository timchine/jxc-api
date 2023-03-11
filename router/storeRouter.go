package router

import (
	"github.com/gin-gonic/gin"
	"github.com/timchine/jxc/api"
	"gorm.io/gorm"
)

func storeRouter(r *gin.RouterGroup, db *gorm.DB) {
	a := api.NewCargoStoreApi(db)
	r.POST("/cargo_store/_search", a.SearchCargoStore())
	r.PUT("/cargo_store/warning", a.SetCargoStoreWarning())
}
